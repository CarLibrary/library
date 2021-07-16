package carlibrary

import (
	"CarLibrary/carlibrary/model"
	"CarLibrary/carlibrary/serializer"
	pb "github.com/CarLibrary/proto/carlibrary"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CarLibraryServiceSever struct {
	pb.UnsafeCarLibraryServiceServer
}

//查看所有品牌
func (c *CarLibraryServiceSever)FindALLCarBand(in *pb.Empty, s pb.CarLibraryService_FindALLCarBandServer) error{

	res,err:=model.FindALLCarBand()
	if err != nil {
		return status.Error(codes.Aborted,err.Error())
	}
	for _,v:=range *res{
		v2:=serializer.BuildCarBand(&v)
		if err=s.Send(v2);err!=nil{
			return status.Error(codes.Unimplemented,err.Error())
		}

	}
	return status.Error(codes.OK,"ok")


}
//查看某品牌的全部车系
func (c *CarLibraryServiceSever)FindAllCarSeries(in *pb.CarSeriesRequest, s pb.CarLibraryService_FindAllCarSeriesServer) error{

	res,err:=model.FindAllCarSeries(in.GetCarBand())
	if err != nil {
		return status.Error(codes.Aborted,err.Error())
	}
	for _,v:=range *res{
		v2:=serializer.BuildCarSeries(&v)
		if err=s.Send(v2);err!=nil{
			return status.Error(codes.Unimplemented,err.Error())
		}
	}
	return status.Error(codes.OK,"ok")

}
//查看某品牌的某车系的全部车型
func (c *CarLibraryServiceSever)FindAllCarModel(in *pb.CarModelRequest,s pb.CarLibraryService_FindAllCarModelServer) error{

	res,err:=model.FindAllCarModel(in.GetCarBand(),in.GetCarSeries())
	if err != nil {
		return status.Error(codes.Aborted,err.Error())
	}
	for _,v:=range *res{
		v2:=serializer.BuildCarModel(&v)
		if err=s.Send(v2);err!=nil{
			return status.Error(codes.Unimplemented,err.Error())
		}
	}
	return status.Error(codes.OK,"ok")
}