package opencc

import (
	"testing"
)

func assertCases(t *testing.T, s2t *OpenCC, cases map[string]string) {
	t.Helper()

	for k, v := range cases {
		str, err := s2t.Convert(k)
		if err != nil {
			t.Error(err)
		}
		if str != v {
			t.Errorf("%s:%s", k, str)
		}
	}
}

func TestConvert_s2t(t *testing.T) {
	cases := map[string]string{
		`我们是工农子弟兵`: `我們是工農子弟兵`,
		`从正数第 x 行到倒数第 y 行，截取多行输出文本的部分内容`:                 `從正數第 x 行到倒數第 y 行，截取多行輸出文本的部分內容`,
		`2017 年中国住房租赁市场租金规模约为 1.3 万亿元`:                   `2017 年中國住房租賃市場租金規模約爲 1.3 萬億元`,
		`香煙（英語：Cigarette），為煙草製品的一種。滑鼠是一種很常見及常用的電腦輸入設備。`:  `香煙（英語：Cigarette），為煙草製品的一種。滑鼠是一種很常見及常用的電腦輸入設備。`,
		`香菸（英語：Cigarette），為菸草製品的一種。記憶體是一種很常見及常用的電腦輸入裝置。`: `香菸（英語：Cigarette），為菸草製品的一種。記憶體是一種很常見及常用的電腦輸入裝置。`,
		`乾隆爷是谁的干爷爷？乾爷爷吗？`:                                `乾隆爺是誰的幹爺爺？乾爺爺嗎？`,
		`2021 年汽车零部件板块市值涨幅跑输乘用车板块，估值相对滞涨，主要由于市场对零部件行业存两大担忧：大宗商品、运费上涨致利润承压；全球芯片紧缺致下游排产低于预期。`: `2021 年汽車零部件板塊市值漲幅跑輸乘用車板塊，估值相對滯漲，主要由於市場對零部件行業存兩大擔憂：大宗商品、運費上漲致利潤承壓；全球芯片緊缺致下游排產低於預期。`,
	}

	s2t, _ := New("s2t")

	assertCases(t, s2t, cases)
}

func TestConvert_s2hk(t *testing.T) {
	cases := map[string]string{}

	s2t, _ := New("s2hk")

	assertCases(t, s2t, cases)
}

func BenchmarkConvert(b *testing.B) {
	s2t, err := New("s2t")
	if err != nil {
		b.Fatal(err)
	}

	// 10621 ns/op in Apple M1
	for n := 0; n < b.N; n++ {
		in := `自然语言处理是人工智能领域中的一个重要方向。`
		out, err := s2t.Convert(in)
		if err != nil {
			b.Fatal(err)
		}
		b.Logf("%s\n%s\n", in, out)
	}
}
