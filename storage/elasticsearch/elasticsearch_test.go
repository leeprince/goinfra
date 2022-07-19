package elasticsearch_test

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/leeprince/goinfra/storage/elasticsearch"
    "os"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/10 下午10:58
 * @Desc:
 */

var (
    localESName = "localESName"
)
var (
    elasticConfs = elasticsearch.ElasticConfs{
        localESName: &elasticsearch.ElasticConf{
            Url:      "https://127.0.0.1:9200",
            Username: "elastic",
            Password: "WmTIU*S0Gw3sS0pEGlVG",
            CACert: `-----BEGIN CERTIFICATE-----
MIIFWjCCA0KgAwIBAgIVAMCEiEeZ0PZp9EQLX8/iVg2DfnRYMA0GCSqGSIb3DQEB
CwUAMDwxOjA4BgNVBAMTMUVsYXN0aWNzZWFyY2ggc2VjdXJpdHkgYXV0by1jb25m
aWd1cmF0aW9uIEhUVFAgQ0EwHhcNMjIwNzA5MTAwMDM4WhcNMjUwNzA4MTAwMDM4
WjA8MTowOAYDVQQDEzFFbGFzdGljc2VhcmNoIHNlY3VyaXR5IGF1dG8tY29uZmln
dXJhdGlvbiBIVFRQIENBMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA
ttOp1ErCKIGQIZWzu+7nD0tZcMcPM/qOhzeCfvKTj0ZL+WnSH5IaTq3L2t0E5fCx
AW3MLy286cdGw75Ce/RfoRp2O3CB/UilnDx/y4TkpLZyM0MRG/xEJTi4W8w0GXO1
bAUtsUbFmcowubyR9aVUJFpxpUAzoC5wYNuGPdeKZEIVMN/G6GDw8zEqoAAZYLNv
r5XlCdC1Bd0qk9b/sSYsUIMaPfePQIM97Dj8zLknfSwzpMTNNqwkkqmMkhoR3Dju
5lUgyPvwIMIGbLHTwd7ckS+VebxffvkfQM4FPIkevE4q9e3yzQ298dqTgLtPPUl5
vaAzN5tVyO+tqz4Lg2nXZWAHLfX0MkXUBzR+95llfCYCRuQnmCvMpNdVlYT2n9nC
YNFRTg6tNLSkzR31XLJ4MRAyuZzXwFD0QZAulDNEsLICX/uFwZlqTqngQWpE1zeb
9oNT8BLNTpoBJ1gPCRPv5QXnXbou6DJpLLRBDJyOGE2Y1y4+WXGAgFgzPy5Wj3Eu
p6TLN9n+4X+eg7Z0MJVpr7QK6F9n7B0t/DmXc5ZQhRyT51RJwZsEmUtfK2MHU4ma
qKnsTtcm7HTq+CC4x/T9v7tekHgukyKJqc4QoWpv9v1cb05SDz8a0n2KBnynJJgM
SDWlK1m2FDFSMmLhiK01TAGEZe7HjQ3BYLC9FLf7AfECAwEAAaNTMFEwHQYDVR0O
BBYEFPONHQtHm6Osta7KtD2MPhBZ9sMKMB8GA1UdIwQYMBaAFPONHQtHm6Osta7K
tD2MPhBZ9sMKMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggIBAD0Z
Gi8UWuRacbXWNDIteEvK/hz9yfdO7gtn19iJtbJaVB6Vymg0zP4YDH74e/Y6yqcz
40YZOu6Op+K3MnoBno3RWft8gp//6qtryNB+yYdXJiXc6T2dvYjbYi98R0/4Wacr
6dw8OqHTVUBLz/27toyOV2qXsxC4WySmWXl/5YQYX4UPk6K27yYsr6Cf7AbWBSbb
5CoOlHgG2vkHg2hCTgg6KiIjRp0hvhLmR28KTbNY94eq3nNNKAEfB0z87LcEiK4Y
831KsHKPV9qwtXcOUHyzyrB7uWVPkZsLEcKAZJKgWlcDUvGZFsyyoYsn6zXyehEr
sM7/at/BYuigzQzHStgmu7oK/C5EBroQhsGHPFsQuubawwfTnlTdLvZQMIbCJFec
gEUipiFq3GqijFs1c01OrUWkqYxi3mpzsNIRiSzn0ppu7twkUaELrIQiZ1e6QH1L
BvGTIeNtyUVPdqsBxQ4b1Xa3s1WD9P6m7SR98UPraBtsj+b8+sd0BSSlSxhSfp6X
KqABVgMxVAjmM+npV845cNdbOTLm4VaSY/k3zu72UQEmu4B2eV1u+ubSFQ7YW9HX
uOshspoYT6Yu4SrXRzwdmwp2rQymHAgN+MU84qWap272kGBfZCg1tuyyoF+H5Hef
J/s0o8YaZwe9ylRV3ICCjEFgRpjqeTVsAPfnfDYx
-----END CERTIFICATE-----
`,
        },
    }
)

func TestMain(m *testing.M) {
    err := elasticsearch.InitElasticClient(elasticConfs)
    fmt.Println("TestMain err", err)
    
    // m.Run 是调用包下面各个Test函数的入口
    os.Exit(m.Run())
}

type Student struct {
    Name         string  `json:"name"`
    Age          int64   `json:"age"`
    AverageScore float64 `json:"average_score"`
}

func TestInitElasticClient(t *testing.T) {
    esClient := elasticsearch.GetElasticClient(localESName)
    if esClient == nil {
        fmt.Println("GetElasticClient nil")
        return
    }
    
    esResp, err := esClient.Info()
    fmt.Println("esResp, err", esResp, err)
    if err != nil {
        fmt.Println("err:", err)
        return
    }
}

func TestDocumentAPIOfCreateUpdate(t *testing.T) {
    esClient := elasticsearch.GetElasticClient(localESName)
    if esClient == nil {
        fmt.Println("GetElasticClient nil")
        return
    }
    
    esResp, err := esClient.Info()
    fmt.Println("esResp, err", esResp, err)
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    
    ctx := context.Background()
    
    studentIndex := "student"
    // --- 插入或者更新文档 ===================================
    // 未指定文档ID：每次指定都会创建
    student := Student{
        Name:         "prince01",
        Age:          18,
        AverageScore: 99.99,
    }
    byteData, err := json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index esResp, err:", esResp, err)
    
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince01",
        Age:          18,
        AverageScore: 99.99,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("1"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince02",
        Age:          18,
        AverageScore: 99.99,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("2"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince03",
        Age:          19,
        AverageScore: 99.99,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("3"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince04",
        Age:          18,
        AverageScore: 80,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("4"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince05",
        Age:          28,
        AverageScore: 99.5,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("5"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // 指定文档ID：不存在则创建，存在则更新
    student = Student{
        Name:         "prince06",
        Age:          20,
        AverageScore: 81,
    }
    byteData, err = json.Marshal(student)
    if err != nil {
        fmt.Println("json marshal err:", err)
        return
    }
    esResp, err = esClient.Index(
        studentIndex,
        bytes.NewReader(byteData),
        esClient.Index.WithDocumentID("6"),
        esClient.Index.WithContext(ctx),
    )
    fmt.Println("esClient.Index WithDocumentID esResp, err:", esResp, err)
    // --- 插入或者更新文档-end ===================================
}

func TestDocumentAPIOfMatch(t *testing.T) {
    esClient := elasticsearch.GetElasticClient(localESName)
    if esClient == nil {
        fmt.Println("GetElasticClient nil")
        return
    }
    
    esResp, err := esClient.Info()
    fmt.Println("esResp, err", esResp, err)
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    
    ctx := context.Background()
    
    studentIndex := "student"
    // --- 搜索文档 +++++++++++++++++++++++++++++
    fmt.Println()
    // var students []Student
    
    // 检查查询
    esResp, err = esClient.Search(
        esClient.Search.WithContext(ctx),
        esClient.Search.WithIndex(studentIndex),
    )
    defer esResp.Body.Close()
    fmt.Println("esClient.Search esResp, err:", esResp, err)
    if err != nil {
        fmt.Println("esClient.Search err")
        return
    }
    fmt.Println("---- 检查查询")
    fmt.Println("esClient.Search esResp.StatusCode:", esResp.StatusCode)
    fmt.Println("esClient.Search esResp.Status():", esResp.Status())
    fmt.Println("esClient.Search esResp.Header:", esResp.Header)
    fmt.Println("esClient.Search esResp.Body:", esResp.Body)
    fmt.Println("esClient.Search esResp.String():", esResp.String())
    fmt.Println("esClient.Search esResp.HasWarnings():", esResp.HasWarnings())
    fmt.Println("esClient.Search esResp.IsError():", esResp.IsError())
    fmt.Println("esClient.Search esResp.Warnings():", esResp.Warnings())
    fmt.Println("----")
    fmt.Println("")
    
    // map多条件查找
    query := map[string]interface{}{
        "from": 0,
        "size": 10,
        "query": map[string]interface{}{
            "bool": map[string]interface{}{
                "must": []interface{}{
                    map[string]interface{}{
                        "range": map[string]interface{}{
                            "age": map[string]interface{}{
                                "from":          18,
                                "include_lower": true,
                                "to":            19,
                                "include_upper": true,
                            },
                        },
                    },
                    map[string]interface{}{
                        "term": map[string]interface{}{
                            "name": "prince01",
                        },
                    },
                },
            },
        },
    }
    var buf bytes.Buffer
    err = json.NewEncoder(&buf).Encode(query)
    
    esResp, err = esClient.Search(
        esClient.Search.WithContext(ctx),
        esClient.Search.WithIndex(studentIndex),
        esClient.Search.WithBody(&buf),
        esClient.Search.WithTrackTotalHits(true),
        esClient.Search.WithTrackScores(false),
        esClient.Search.WithFrom(0),
        esClient.Search.WithSize(5),
    )
    defer esResp.Body.Close()
    fmt.Println("esClient.Search esResp, err:", esResp, err)
    if err != nil {
        fmt.Println("esClient.Search err")
        return
    }
    fmt.Println("---- map多条件查找")
    fmt.Println("esClient.Search esResp.StatusCode:", esResp.StatusCode)
    fmt.Println("esClient.Search esResp.Status():", esResp.Status())
    fmt.Println("esClient.Search esResp.Header:", esResp.Header)
    fmt.Println("esClient.Search esResp.Body:", esResp.Body)
    fmt.Println("esClient.Search esResp.String():", esResp.String())
    fmt.Println("esClient.Search esResp.HasWarnings():", esResp.HasWarnings())
    fmt.Println("esClient.Search esResp.IsError():", esResp.IsError())
    fmt.Println("esClient.Search esResp.Warnings():", esResp.Warnings())
    fmt.Println("----")
    fmt.Println("")
    
    // --- 搜索文档-end +++++++++++++++++++++++++++++
}
