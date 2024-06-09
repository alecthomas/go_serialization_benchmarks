using Go = import "/github.com/glycerine/go-capnproto/go.capnp";
$Go.package("capnproto");
$Go.import("github.com/alecthomas/go_serialization_benchmarks/internal/capnproto");

@0x99ea7c74456111bd;

struct CapnpA {
  name     @0   :Text;
  birthDay @1   :Int64;
  phone    @2   :Text;
  siblings @3   :Int32;
  spouse   @4   :Bool;
  money    @5   :Float64;
}
