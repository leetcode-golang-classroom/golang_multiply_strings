# golang_multiply_strings

Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

**Note:** You must not use any built-in BigInteger library or convert the inputs to integer directly.

## Examples

**Example 1:**

```
Input: num1 = "2", num2 = "3"
Output: "6"

```

**Example 2:**

```
Input: num1 = "123", num2 = "456"
Output: "56088"

```

**Constraints:**

- `1 <= num1.length, num2.length <= 200`
- `num1` and `num2` consist of digits only.
- Both `num1` and `num2` do not contain any leading zero, except the number `0` itself.

## 解析

給定兩個數字字串 num1, num2

要求寫一個演算來來做兩個數字相乘並且回傳相乘後的字串

題目特別要求不要直接使用字串轉數字的方式去做

也就是要採用一個 digit 一個 digit 的方式去運算

因為把兩數相乘以位元的角度是把其中一個數值固定

變動另一個數字在另一個數字的每個位元當作起始點

可以發現有以下關係

![](https://i.imgur.com/NEem4t0.png)


而新位元的每個位元 digits 剛好會是原本兩個數字個取一個位元相乘 + carry 上一個位元運算的進位 % 10

而下個carry原本兩個數字個取一個位元相乘 + carry 上一個位元運算的進位 / 10

如下圖：

![](https://i.imgur.com/9Z4PvqO.png)

![](https://i.imgur.com/98rqrFA.png)

透過這樣的方式

只需要 採用 O(n*m) 的時間複雜度

還有 O(n+m) 的空間複雜度

就可以做完整個計算

## 程式碼
```go
package sol

func multiply(num1 string, num2 string) string {
	lenNum1, lenNum2 := len(num1), len(num2)
	fixLen := lenNum1 + lenNum2 - 1
	carry := 0
	tempDigits := make([]int, fixLen)
	for digit := 0; digit < fixLen; digit++ {
		sum := 0
		// digit = lenNum1 - 1 - pos + lenNum2 - 1 - c_pos
		// c_pos = lenNum1 - 1 - pos + lenNum2 - 1 - digit
		for pos := lenNum1 - 1; pos >= 0; pos-- {
			c_pos := lenNum1 - 1 - pos + lenNum2 - 1 - digit
			if c_pos >= 0 && c_pos < lenNum2 {
				sum += int(num1[pos]-'0') * int(num2[c_pos]-'0')
			}
		}
		sum += carry
		tempDigits[fixLen-1-digit] = sum % 10
		carry = sum / 10
	}
	res := []byte{}
	if carry > 0 {
		ch := byte(carry + '0')
		res = append(res, ch)
	}
	for idx, val := range tempDigits {
		ch := byte(val + '0')
		res = append(res, ch)
		if ch-'0' == 0 && idx == 0 && carry == 0 {
			break
		}
	}
	return string(res)
}

```
## 困難點

1. 需要看兩數相乘 digit 之前的關係
2. 需要理解要從最低位元運算到最高位元

## Solve Point

- [x]  初始化 lenNum1 := len(num1),lenNum2 := len(num2) fixLen := lenNum1 + lenNum2 - 1
- [x]  建立 tempDigits 為長度為 fixLen 的整數陣列, carry = 0
- [x]  從 digit := 0 開始計算到 digit = fixLen - 1
- [x]  每次 符合 digit = len(num1) - 1 - pos + len(num2) - 1 - c_pos, pos 與 c_pos 來做運算
- [x]  紀錄 sum += (num1[pos] - ‘0’) * (num2[c_pos]-’0’)
- [x]  sum += carry
- [x]  更新 tempDigits[digit] = sum %10
- [x]  更新 carry = sum / 10
- [x]  新增 result = [] byte array
- [x]  if carry > 0 , 把 carry 加入 result
- [x]  把 tempDigits 所有元素加入 result
- [x]  把 byteArray 轉換成 string
