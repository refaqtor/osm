package osm

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestFeature_String(t *testing.T) {
	cases := []struct {
		name     string
		id       FeatureID
		expected string
	}{
		{
			name:     "node",
			id:       NodeID(1).FeatureID(),
			expected: "node/1",
		},
		{
			name:     "way",
			id:       WayID(3).FeatureID(),
			expected: "way/3",
		},
		{
			name:     "relation",
			id:       RelationID(1000).FeatureID(),
			expected: "relation/1000",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if v := tc.id.String(); v != tc.expected {
				t.Errorf("incorrect string: %v", v)
			}
		})
	}
}

func TestFeatureIDsSort(t *testing.T) {
	ids := FeatureIDs{
		RelationID(1).FeatureID(),
		ChangesetID(1).FeatureID(),
		NodeID(1).FeatureID(),
		WayID(2).FeatureID(),
		WayID(1).FeatureID(),
		ChangesetID(3).FeatureID(),
		ChangesetID(1).FeatureID(),
	}

	expected := FeatureIDs{
		NodeID(1).FeatureID(),
		WayID(1).FeatureID(),
		WayID(2).FeatureID(),
		RelationID(1).FeatureID(),
		ChangesetID(1).FeatureID(),
		ChangesetID(1).FeatureID(),
		ChangesetID(3).FeatureID(),
	}

	ids.Sort()
	if !reflect.DeepEqual(ids, expected) {
		t.Errorf("not sorted correctly")
		for i := range ids {
			t.Logf("%d: %v", i, ids[i])
		}
	}
}

func BenchmarkFeatureIDsSort(b *testing.B) {
	rand.Seed(1024)

	tests := make([]FeatureIDs, b.N)
	for i := range tests {
		ids := make(FeatureIDs, 10000)

		for j := range ids {
			switch rand.Intn(4) {
			case 0:
				ids[j] = NodeID(rand.Int63n(int64(len(ids) / 10))).FeatureID()
			case 1:
				ids[j] = WayID(rand.Int63n(int64(len(ids) / 10))).FeatureID()
			case 2:
				ids[j] = RelationID(rand.Int63n(int64(len(ids) / 10))).FeatureID()
			case 3:
				ids[j] = ChangesetID(rand.Int63n(int64(len(ids) / 10))).FeatureID()
			}
		}
		tests[i] = ids
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tests[n].Sort()
	}
}
