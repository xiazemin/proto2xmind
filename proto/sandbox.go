package gen

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emicklei/proto"
	xmlXmind "github.com/xiazemin/xmind-go/xxml"
)

type SandBox struct {
	num                   int64 //start num
	packageId             string
	messageTypeToParentId map[string][]string
	meaasgeTable          map[string]*proto.Message
	*Xmind
}

func NewSandBox(num int64, xmind *Xmind) *SandBox {
	return &SandBox{
		num:                   num,
		packageId:             "",
		messageTypeToParentId: make(map[string][]string),
		meaasgeTable:          make(map[string]*proto.Message),
		Xmind:                 xmind,
	}
}

func (s *SandBox) Proto2Xmind(src string) int64 {
	reader, err := os.Open(src)

	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()
	if len(s.sheet) == 0 {
		s.sheet = append(s.sheet, &xmlXmind.XmindNode{
			NodeID:       "1",
			TopicContent: "root.proto",
		})
	}

	proto.Walk(definition,
		proto.WithPackage(s.handlePackage),
		proto.WithService(s.handleService),
	)
	//解决嵌套问题
	proto.Walk(definition,
		proto.WithMessage(s.preHandleMessage),
	)
	//解决声明顺序的问题
	proto.Walk(definition,
		proto.WithMessage(s.handleMessage),
	)
	return s.num
}

func (s *SandBox) handlePackage(p *proto.Package) {
	s.num++
	s.sheet = append(s.sheet, &xmlXmind.XmindNode{
		NodeID:       strconv.FormatInt(s.num, 10),
		TopicContent: p.Name, //package
		ParentID:     "1",
	})
	s.packageId = strconv.FormatInt(s.num, 10)
	//s.Accept(root)
}
func (s *SandBox) handleService(srv *proto.Service) {
	s.num++
	s.sheet = append(s.sheet, &xmlXmind.XmindNode{
		NodeID:       strconv.FormatInt(s.num, 10),
		TopicContent: srv.Name,    //service
		ParentID:     s.packageId, //package
	})
	root := serviceLister{
		parentId: strconv.FormatInt(s.num, 10),
		SandBox:  s,
	}
	//s.Accept(root)
	for _, e := range srv.Elements {
		e.Accept(root)
		//fmt.Println(i)
	}
}

func (s *SandBox) preHandleMessage(m *proto.Message) {
	//fmt.Println("preMsg:", m.Name)
	s.meaasgeTable[m.Name] = m
}

func (s *SandBox) handleMessage(m *proto.Message) {
	lister := new(messageLister)
	lister.SandBox = s
	if refers := s.messageTypeToParentId[m.Name]; len(refers) > 0 {
		// num++
		// lister.parentId = strconv.FormatInt(num, 10)
		// s.sheet = append(s.sheet, &xmlXmind.XmindNode{
		// 	NodeID:       strconv.FormatInt(num, 10),
		// 	TopicContent: m.Name, //response
		// 	ParentID:     refer,  //函数参数或者返回值
		// })
		lister.parentIds = refers
		for _, each := range m.Elements {
			each.Accept(lister)
		}
	}
	//fmt.Println(m.Name)
}
