package fakertest

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/11 14:00
 * @Desc:
 */

func TestFaker_test(t *testing.T) {
	fmt.Println("faker.UUIDHyphenated", faker.UUIDHyphenated())
	fmt.Println("faker.Name", faker.Name())
	fmt.Println("faker.Word", faker.Word())
}
