package scannerapi

import (
	"context"
	"crypto/rsa"
	"testing"

	"github.com/fox-one/pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

const (
	appID      = "d5110a2d-8a8d-4ab1-874d-155577ddad90"
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAoudQbjJswXP6Gw1v5AkoezaQ0cWVIFakdJHqGvLlbT72e9z5
AVdtawNtVeS7QZSv5+hdbXxHYqQmmokHPoIHoYCntvgT+vU1geZ3mNl/oZxtTc13
rrVSZso9slGWLi/0oz/y1jJmOaxcLNmHnKHYCQ2ouHribH4/v7+Rq+jhoVboZUCT
Dkclqoio4ae4xE1cGx5yC4B8IYTRW1Ahtv7pNaoJbxhMjnR/qK5scXB3sWa9ypNC
JEOXi0OJUvLAZ7GY7TD+gKYE5JduX4lkOgpN/SBeQXeK78BUgzXy7MdMQC31J31c
+HHWW8VlBh6DEa+HC/hf9+wHfbYp/V1mFCfTEQIDAQABAoIBADNyK8NxZ0k88DqE
5tY0UjV/SQMGjA73bd5minFSTkRPAxag9X6H9sU6AtiCcOTIKTlq975w5jZvssVR
CxmhhPlrIQw2klDKCcjpWNHLtnP6a8aLYgWpu8k8i3l6LJyPhonb0zv0FLaYCVAc
rqg3sHtDdgo0vLcYs9dBj6aj2EXUJGf1dEjM7rp0k3HUK84pBgakXHgLrQo+vaBK
WaCH+AmAwvLPOtl90osSzgn4sSxUmwA9gbe6sT3fBULX8KgsJ9BYLGxLwKs/+UMW
tYqVa7PUmDgBky1DER+q6rwQL8EnQvMtvm/XjZvF97dToPrTKSTfw10WWJ5kJRX4
7iFkPnECgYEA0xa9iT38+5tXasRI7Sx/3iiuZ3LtdKxqi/spvP6Mm1fzb0PMjzRW
zE3R4rOZhEIZzet9jl7CKB5IphtPnQZCzzNUMieUkhCz9sjEHyp/PEf2ZJRjab6R
l4yjOk+DLoMh0Kjrypoi6VoreRAs6KcLMdjckF8jmUoh0SFc37Zlo20CgYEAxZAX
/peDebXlrpfHMZVLPqWb/84kdkvnhdeVs9vDy80BXezhRaLvoDixs/jORCEk2Fwn
OsNirI70SxS98qhuzO+OZPZmIHNUISds/B4QmcmkrLJc21WX/jmZlLhLDrGRoHzU
EmqzvU69Pco9chzp673aKn999fEGbq0/S8VFA7UCgYA7NpXUed8NdXYALS+H2IPF
/gNuYX7ay8LXoY0DxyQKL5HKFGq9gSEYDnxh65+UyjYS3YxK86njNxuJ/P9IPQlg
tjVYHGJA1w6km4zocTNf5Y/ohvecIikVKg5fvQ4/bel+buQ14ScJ7pFjVHZEjRdG
1F6K4SVLStBRqdozpya7wQKBgCQzkEIFr4DY9Pp2M3kBe2VCYYCRXJsUs5sR8wuf
JJSuvyZf5rOk9LXuTZnF6L4ROLgwwruA3A70rw0nAtt5Q7xn1Tbo87PUjkD5AX59
X9teWL+Qe3VxjMV39o2K664Imb/Hv/CT/tWcnQ66hWHorHCXPGf3LrSnct9K/cM7
pbbVAoGBAJynDrjsGh2GSwJhVCyOnNOvcV+PXDb+uHSnNcI97mz0ou3lfLvCrWE/
0xP+YOK3IbDFJba9kN9d1YsSY+ernnwC9RmKFKfz7yS09TtUfRLfdZ2PGGB9veR0
/rPzFK1R1kHEEAKa1Uooy63d1n8Ck2/E/uvpSSsrpZVFcrPx3Nv5
-----END RSA PRIVATE KEY-----`
)

func TestQueryOrders(t *testing.T) {
	UseEndpoint("http://localhost:8082")

	key, _, err := encrypt.ParsePrivatePem(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	app := App{
		AppID:      appID,
		PrivateKey: key.(*rsa.PrivateKey),
	}
	ctx := WithApp(context.Background(), &app)
	_, err = ScanAssetBalances(ctx, "94d6c6a3-3d3b-35f9-ae76-bbaa94c0caa9", 0)
	assert.Nil(t, err)
}
