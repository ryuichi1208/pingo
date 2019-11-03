import (
  "google.golang.org/grpc"
  pb "google.golang.org/grpc/examples/route_guide/routeguide"
)

func printFeature(client pb.RouteGuideClient, point *pb.Point) {
  log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  feature, err := client.GetFeature(ctx, point)
  if err != nil {
    log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
  }
  log.Println(feature)
}

func main {
  var opts []grpc.DialOption
  opts = append(opts, grpc.WithInsecure()) // Connection uses plain TCP, TLS also exists

  serverAddr := "127.0.0.1:10000"
  conn, err := grpc.Dial(*serverAddr, opts...)

  if err != nil {
    log.Fatalf("fail to dial: %v", err)
  }
  defer conn.Close()

  client := pb.NewRouteGuideClient(conn)
  printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
}
