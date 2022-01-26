package handlers

import (
	proto "trussdemo"
	"trussdemo/prototool"
)

type Event struct {
	Name               string
	EndTime            *int32
	RestrictionRegions []string
}

// I think generate can help us out here
func modelToProto(e *Event) *proto.Event {
	if e == nil {
		return nil
	}
	pb := &proto.Event{
		Name:    e.Name,
		EndTime: prototool.OptionInt32(e.EndTime),
	}
	pb.RestrictionRegions = make([]proto.Event_Region, 0, len(e.RestrictionRegions))
	for _, region := range e.RestrictionRegions {
		pb.RestrictionRegions = append(pb.RestrictionRegions, proto.Event_Region(proto.Event_Region_value[region]))
	}
	return pb
}

func protoToModel(event *proto.Event) *Event {
	if event == nil {
		return nil
	}
	return &Event{
		Name:    event.GetName(),
		EndTime: prototool.PInt32(event.GetEndTime()),
	}
}
