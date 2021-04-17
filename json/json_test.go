package json

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 12:52
 * @Desc:
 */

type User struct {
	Id        int    `json:"id"`
	StartTime int    `json:"startTime"`
	Role      string `json:"role"`
	Relation  string `json:"relation"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	Status    string `json:"status"`
}

func TestJsonEncode(t *testing.T) {
	tests := []struct {
		user   *User
		expect string
	}{
		{
			&User{1, 123, "STUDENT", "FATHER", "wei", "chuyi ", "status"},
			`{"id":1,"startTime":123,"role":"STUDENT","relation":"FATHER","name":"wei","class":"chuyi ","status":"status"}`},
	}

	for _, test := range tests {
		if actual := JsonEncode(test.user); string(actual) != test.expect {
			t.Errorf("JsonEncode(%v): expect %s, actual %s",
				test.user, test.expect, actual)
		}
	}
}

func TestJsonEncodeToString(t *testing.T) {
	tests := []struct {
		user   *User
		expect string
	}{
		{
			&User{1, 123, "STUDENT", "FATHER", "wei", "chuyi ", "status"},
			`{"id":1,"startTime":123,"role":"STUDENT","relation":"FATHER","name":"wei","class":"chuyi ","status":"status"}`},
	}

	for _, test := range tests {
		if actual := JsonEncodeToString(test.user); actual != test.expect {
			t.Errorf("JsonEncodeToString(%v): expect %s, actual %s",
				test.user, test.expect, actual)
		}
	}
}

func TestJsonDecode(t *testing.T) {

	tests := []struct {
		user   string
		expect User
	}{
		{

			`{"id":1,"startTime":123,"role":"STUDENT","relation":"FATHER","name":"wei","class":"chuyi ","status":"status"}`,
			User{1, 123, "STUDENT", "FATHER", "wei", "chuyi ", "status"},
		},
	}

	for _, test := range tests {
		var actual User
		if err := JsonDecode([]byte(test.user), &actual); actual != test.expect {
			if err != nil {
				t.Error(err)
				continue
			}
			t.Errorf("JsonDecode(%s): expect %v, actual %v",
				test.user, test.expect, actual)
		}
	}
}
