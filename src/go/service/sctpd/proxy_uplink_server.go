// Copyright 2021 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sctpd

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/magma/magma/log"
	pb "github.com/magma/magma/protos/magma/sctpd"
)

// ProxyUplinkServer handles SctpdUplinkServer RPCs by calling out to a
// provided SctpdUplinkClient.
type ProxyUplinkServer struct {
	client pb.SctpdUplinkClient

	log.Logger
	*pb.UnimplementedSctpdUplinkServer
}

// NewProxyUplinkServer creates a new ProxyUplinkServer with the provided
// logger and client connection.
func NewProxyUplinkServer(logger log.Logger, cc *grpc.ClientConn) *ProxyUplinkServer {
	return &ProxyUplinkServer{
		Logger: logger,
		client: pb.NewSctpdUplinkClient(cc),
	}
}

// SendUl proxies calls to SctpdUplink.SendUl.
func (p *ProxyUplinkServer) SendUl(ctx context.Context, req *pb.SendUlReq) (*pb.SendUlRes, error) {
	p.Logger.
		With("assoc_id", req.GetAssocId()).
		With("stream", req.GetStream()).
		With("ppid", req.GetPpid()).
		With("payload_size", len(req.GetPayload())).
		Debug().Print("SendUl")
	return p.client.SendUl(ctx, req)
}

// NewAssoc proxies calls to SctpdUplink.NewAssoc.
func (p *ProxyUplinkServer) NewAssoc(ctx context.Context, req *pb.NewAssocReq) (*pb.NewAssocRes, error) {
	p.Logger.
		With("assoc_id", req.GetAssocId()).
		With("instreams", req.GetInstreams()).
		With("outstreams", req.GetOutstreams()).
		With("ppid", req.GetPpid()).
		With("ran_cp_ipaddr", net.IP(req.GetRanCpIpaddr()).String()).
		Debug().Print("NewAssoc")
	return p.client.NewAssoc(ctx, req)
}

// CloseAssoc proxies calls to SctpdUplink.CloseAssoc.
func (p *ProxyUplinkServer) CloseAssoc(ctx context.Context, req *pb.CloseAssocReq) (*pb.CloseAssocRes, error) {
	p.Logger.
		With("assoc_id", req.GetAssocId()).
		With("is_reset", req.GetIsReset()).
		With("ppid", req.GetPpid()).
		Debug().Print("CloseAssoc")
	return p.client.CloseAssoc(ctx, req)
}
