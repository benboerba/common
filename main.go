package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	questionid string = "311745535"
	cookie     string = ``
	prefix     string = "https://www.zhihu.com/api/v4/questions/"
	postfix    string = "/answers?include=data%5B%2A%5D.is_normal%2Cadmin_closed_comment%2Creward_info%2Cis_collapsed%2Cannotation_action%2Cannotation_detail%2Ccollapse_reason%2Cis_sticky%2Ccollapsed_by%2Csuggest_edit%2Ccomment_count%2Ccan_comment%2Ccontent%2Ceditable_content%2Cvoteup_count%2Creshipment_settings%2Ccomment_permission%2Ccreated_time%2Cupdated_time%2Creview_info%2Crelevant_info%2Cquestion%2Cexcerpt%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%2Cis_labeled%2Cis_recognized%2Cpaid_info%3Bdata%5B%2A%5D.mark_infos%5B%2A%5D.url%3Bdata%5B%2A%5D.author.follower_count%2Cbadge%5B%2A%5D.topics&limit=20&offset="
	ppos       string = "&sort_by=created"
	myfilepath string = ``
)

var url string = "https://www.zhihu.com/api/v4/questions/311745535/answers?include=data%5B%2A%5D.is_normal%2Cadmin_closed_comment%2Creward_info%2Cis_collapsed%2Cannotation_action%2Cannotation_detail%2Ccollapse_reason%2Cis_sticky%2Ccollapsed_by%2Csuggest_edit%2Ccomment_count%2Ccan_comment%2Ccontent%2Ceditable_content%2Cvoteup_count%2Creshipment_settings%2Ccomment_permission%2Ccreated_time%2Cupdated_time%2Creview_info%2Crelevant_info%2Cquestion%2Cexcerpt%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%2Cis_labeled%2Cis_recognized%2Cpaid_info%3Bdata%5B%2A%5D.mark_infos%5B%2A%5D.url%3Bdata%5B%2A%5D.author.follower_count%2Cbadge%5B%2A%5D.topics&limit=20&offset=0&sort_by=created"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resp_byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	tmp := make(map[string]interface{})
	err = json.Unmarshal(resp_byte, &tmp)
	if err != nil {
		fmt.Println(err)
	}

	tmp2 := tmp["data"].([]interface{})
	tmp_len := len(tmp2) // 每个分段的长度
	fmt.Println(tmp_len)
	for _, val := range tmp2 {
		tmp3 := val.(map[string]interface{})
		tmp4 := tmp3["created_time"]
		//fmt.Println(tmp4) // 打出来是个科学计数法
		fmt.Printf("%.f\n", tmp4)
	}

	return
}