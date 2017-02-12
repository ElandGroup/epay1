package mapHelper

func ConvMap(in map[string]string) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for i, v := range in {
		out[i] = v
	}
	return
}
