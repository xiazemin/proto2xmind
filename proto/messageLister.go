package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emicklei/proto"
	xmlXmind "github.com/xiazemin/xmind-go/xxml"
)

type messageLister struct {
	proto.NoopVisitor
	parentIds []string
	depth     int64
	*SandBox
}

//VisitProto(p *Proto)
func (l messageLister) VisitMessage(m *proto.Message) {

}
func (l messageLister) VisitService(v *proto.Service) {

}
func (l messageLister) VisitSyntax(s *proto.Syntax) {

}
func (l messageLister) VisitPackage(p *proto.Package) {

}
func (l messageLister) VisitOption(o *proto.Option) {
	fmt.Println(o.Name)

}
func (l messageLister) VisitImport(i *proto.Import) {

}
func (l messageLister) VisitNormalField(i *proto.NormalField) {
	//fmt.Println(i.Name, l.depth)
	l.num++

	fieldType := i.Type
	if i.Repeated {
		fieldType = "[]" + fieldType
	}

	comment := " "
	if i != nil && i.Comment != nil {
		comment = comment + strings.Join(i.Comment.Lines, " ")
	}
	for _, pid := range l.parentIds {
		l.sheet = append(l.sheet, &xmlXmind.XmindNode{
			NodeID:       strconv.FormatInt(l.num, 10),
			TopicContent: strconv.FormatInt(int64(i.Sequence), 10) + ":" + i.Name + " " + fieldType + comment, //package
			ParentID:     pid,
		})

		if msgType := l.meaasgeTable[i.Type]; msgType != nil {
			parentIds := l.parentIds
			l.depth++
			//fmt.Println("subMsg:", i.Type, msgType.Name)
			l.parentIds = []string{strconv.FormatInt(l.num, 10)}
			//msgType.Accept(l)
			for _, each := range msgType.Elements {
				each.Accept(l)
			}
			l.depth--
			l.parentIds = parentIds
		}
	}
	//i.Comment.Accept(l)
}
func (l messageLister) VisitEnumField(i *proto.EnumField) {

}
func (l messageLister) VisitEnum(e *proto.Enum) {

}
func (l messageLister) VisitComment(e *proto.Comment) {
	if e != nil {
		//fmt.Println(e.Lines)
	}
}
func (l messageLister) VisitOneof(o *proto.Oneof) {
	//fmt.Println("oneof:", o.Name)
	for _, e := range o.Elements {
		e.Accept(l)
	}
}

func (l messageLister) VisitOneofField(o *proto.OneOfField) {
	//fmt.Println("oneofField:", o.Field.Name)
	l.num++
	comment := " "
	if o != nil && o.Comment != nil {
		comment = comment + strings.Join(o.Comment.Lines, " ")
	}
	for _, pid := range l.parentIds {
		l.sheet = append(l.sheet, &xmlXmind.XmindNode{
			NodeID:       strconv.FormatInt(l.num, 10),
			TopicContent: strconv.FormatInt(int64(o.Sequence), 10) + ":" + "oneof:" + o.Field.Name + " " + o.Field.Type + comment, //package
			ParentID:     pid,
		})

		if msgType := l.meaasgeTable[o.Field.Type]; msgType != nil {
			parentIds := l.parentIds
			l.depth++
			//fmt.Println("subMsg:", o.Field.Type, msgType.Name)
			l.parentIds = []string{strconv.FormatInt(l.num, 10)}
			//msgType.Accept(l)
			for _, each := range msgType.Elements {
				each.Accept(l)
			}
			l.depth--
			l.parentIds = parentIds
		}
	}
}

func (l messageLister) VisitReserved(r *proto.Reserved) {

}
func (l messageLister) VisitRPC(r *proto.RPC) {

}
func (l messageLister) VisitMapField(f *proto.MapField) {

}

// proto2
func (l messageLister) VisitGroup(g *proto.Group) {

}
func (l messageLister) VisitExtensions(e *proto.Extensions) {

}
