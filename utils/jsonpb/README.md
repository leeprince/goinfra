# json protobuff 序列化

基于`github.com/golang/protobuf@1.x/jsonpb`库和`github.com/grpc-ecosystem/grpc-gateway@v1.xx/runtime`库开发的JSONPb Marshaler包,
提供一个新的JSONPb器，用于替换 `runtime.JSONPb`。

**替换原因**：对于在Proto文件中定义的`int64`/`uint64`/`Enum`字段，默认的`runtime.JSONPb`会将其序列化为`string`类型。

当gPRC接口是用于提供给前端JS调用时，将高精度数值转成`string`会更安全，这是因为JS本身的基础数据类型精度位数就有限。

但除此之外，通过开放平台对外开放的接口，对方可能是通过PHP/Java/Python/Golang进行调用的。

在这种情况下我们希望能够自主决定序列化的行为，而不是被`protobuf`库写死在代码中:[https://github.com/golang/protobuf/blob/master/jsonpb/encode.go#L547-L550](https://github.com/golang/protobuf/blob/master/jsonpb/encode.go#L547-L550)


例如：
```
// Proto文件定义
message Department {
    int64 id = 1;
    string name = 2;
    int32 type = 3;
}
```

```golang
type Department struct {
    ID int64 
    Name string
    Type int32
}
func (s *GRPCServer) GetDepartment(ctx context.Context, req *pb.GetDepartmentReq) (*pb.GetDepartmentRsp, error) {
    rsp := &pb.GetDepartmentRsp{
        Department: &pb.Department{
            Id: 1379,
            Name: "some department",
            Type: 1,
        },
    }
    return rsp, nil
}
```

通过`grpc-echosystem/grpc-gateway`转换后会序列化为:
```json
{
    "id": "1379",
    "name": "some department",
    "type": 1
}
```

而`gcjsonpb.JSONPb`就是解决这个问题的。


