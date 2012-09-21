package topology

import (
	"strconv"
	"testing"
  _ "fmt"
)

func TestXYZ(t *testing.T) {
	topo := NewTopology("topo","/tmp","test",234,5)
	for i := 0; i < 5; i++ {
		dc := NewDataCenter("dc" + strconv.Itoa(i))
		dc.activeVolumeCount = i
		dc.maxVolumeCount = 5
		topo.LinkChildNode(dc)
	}
	nl := NewNodeList(topo.Children(),nil)

  picked, ret := nl.RandomlyPickN(1)
  if !ret || len(picked)!=1 {
    t.Errorf("need to randomly pick 1 node")
  }

	picked, ret = nl.RandomlyPickN(4)
	if !ret || len(picked)!=4 {
	  t.Errorf("need to randomly pick 4 nodes")
	}

  picked, ret = nl.RandomlyPickN(5)
  if !ret || len(picked)!=5 {
    t.Errorf("need to randomly pick 5 nodes")
  }

  picked, ret = nl.RandomlyPickN(6)
  if ret || len(picked)!=0 {
    t.Errorf("can not randomly pick 6 nodes:", ret, picked)
  }

}
