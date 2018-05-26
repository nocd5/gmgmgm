package gmgmgm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	list := []string{
		"apple",
		"orange",
		"grape",
		"banana",
		"cherry",
		"pineapple",
		"lemon",
		"melon",
		"kiwi",
	}

	assert.Equal(t, Match("ap", list, false, false),
		[]string{"apple", "grape", "pineapple"})
	assert.Equal(t, Match("ap", list, false, true),
		[]string{"apple"})
	assert.Equal(t, Match("an", list, false, false),
		[]string{"banana", "orange"})
	assert.Empty(t, Match("an", list, false, true))
}

func Test2(t *testing.T) {
	list := []string{
		"Apple",
		"Grape",
		"Banana",
		"Cherry",
		"Pineapple",
		"Lemon",
		"Melon",
		"Kiwi",
	}

	assert.Equal(t, Match("ap", list, false, false),
		[]string{"Grape", "Pineapple"})
	assert.Equal(t, Match("ap", list, true, false),
		[]string{"Apple", "Grape", "Pineapple"})
	assert.Equal(t, Match("ap", list, true, true),
		[]string{"Apple"})
	assert.Empty(t, Match("ap", list, false, true))
}

func Test3(t *testing.T) {
	list := []string{
		"りんご",
		"リンゴ",
		"林檎",
		"ぶどう",
		"ブドウ",
		"葡萄",
		"武道",
		"バナナ",
		"さくらんぼ",
		"パイナップル",
		"レモン",
		"メロン",
		"キウイ",
		"かき",
		"柿",
	}

	assert.Equal(t, Match("ri", list, false, false),
		[]string{"りんご", "リンゴ", "林檎"})
	assert.Equal(t, Match("budou", list, false, false),
		[]string{"ぶどう", "ブドウ", "武道", "葡萄"})
	assert.Equal(t, Match("ki", list, false, false),
		[]string{"かき", "キウイ", "林檎"}) // 檎:キン
	assert.Equal(t, Match("ki", list, false, true),
		[]string{"キウイ"})
}

func Test4(t *testing.T) {
	list := []string{
		"test",
		"試験",
		"テスト",
		"computer",
		"コンピューター",
		"コンピュータ",
	}

	assert.Equal(t, Match("test", list, false, false),
		[]string{"test", "テスト", "試験"})
	assert.Equal(t, Match("comp", list, false, false),
		[]string{"computer", "コンピュータ", "コンピューター"})
	assert.Equal(t, Match("te", list, false, false),
		[]string{"computer", "test", "テスト"})
	assert.Equal(t, Match("te", list, true, false),
		[]string{"computer", "test", "テスト"})
	assert.Equal(t, Match("te", list, true, true),
		[]string{"test", "テスト"})
}
