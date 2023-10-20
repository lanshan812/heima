package patent

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"patent/service"
	err "patent/util"
	"patent/util/log"
	"regexp"

	"github.com/google/go-tika/tika"
)

type SystemService struct {
	service.Service
}

func NewSystemService(account string) *SystemService {
	return &SystemService{service.Service{Account: account}}
}

func (s *SystemService) ParsingPdf(param *multipart.FileHeader) (*GetCarScheduleListResp, *err.Error) {
	//判断文件是否接收成功
	f, e := param.Open()
	if e != nil {
		erInfo := "文件接受失败"
		log.Logger.Error(erInfo)
		return nil, err.New(400, "InvalidParam", erInfo, "ParsingPdf", e)
	}
	//接收pdf解析的数据
	resumeData, e := ProcessTikaData(f)
	if e != nil {
		erInfo := "pdf转word失败"
		log.Logger.Error(erInfo)
		return nil, err.New(500, "StatusInternalServerError", erInfo, "ParsingPdf", e)
	}
	fmt.Println(resumeData)

	//正则匹配保存到数据库中
	//正则匹配
	/* (10)授权公告号
	（54）xx名称
	（71）申请人
	（12）专利类型
	  技术领域
	   （22）申请日
	   （45）授权公布日
	   （72）发明人
	   （74）专利代理机构
	   （57）摘要

	*/
	// reg1, _ := regexp.Compile(`申请公布号(.+)`)
	// reg2 := regexp.MustCompile(`54([\s\S]*?)\(`)
	// reg3 := regexp.MustCompile(`\(71\)申请人(.+)`)
	// reg4 := regexp.MustCompile(`\(12\)专利类型(.+)`)
	// reg5 := regexp.MustCompile(`技术领域(.+)`)
	// reg6 := regexp.MustCompile(`\(22\)申请日(.+)`)
	// reg7 := regexp.MustCompile(`\(45\)授权公布日(.+)`)
	// reg8 := regexp.MustCompile(`\(72\)发明人(.+)`)
	// reg9 := regexp.MustCompile(`专利代理机构 (.+)`)
	// reg10 := regexp.MustCompile(`\(57\)摘要(.+)`)
	// str1 := reg1.FindAllStringSubmatch(resumeData, -1)
	// str2 := reg2.FindAllStringSubmatch(resumeData, -1)
	// str3 := reg3.FindAllStringSubmatch(resumeData, -1)
	// str4 := reg4.FindAllStringSubmatch(resumeData, -1)
	// str5 := reg5.FindAllStringSubmatch(resumeData, -1)
	// str6 := reg6.FindAllStringSubmatch(resumeData, -1)

	//找到所有的成切片
	// var patentsave = model.PatentInfo{
	// 	PatentId:        str1[0][0],
	// 	PatentName:      str2[0][0],
	// 	Abstract:        str3[0][0],
	// 	PatentType:      str4[0][0],
	// 	PatentField:     str5[0][0],
	// 	ApplicationDate: str6[0][0],
	// }
	// fmt.Println(patentsave)
	return nil, nil
}

// 调用Tika库生成字符串
func ProcessTikaData(f io.Reader) (resumeContent []string, err error) {
	//模拟客户端
	client := tika.NewClient(nil, "http://localhost:9998")
	//转成str的html页面
	context1, err := client.Parse(context.TODO(), f)
	fmt.Println(context1)
	reg := regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	//找到所有的成切片
	res := reg.FindAllStringSubmatch(context1, -1)
	//遍历添加到resumeContent
	for _, match := range res {
		resumeContent = append(resumeContent, match...)
	}
	return resumeContent, err
}
