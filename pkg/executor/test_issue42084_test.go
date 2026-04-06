package executor_test

import (
	"testing"
	"github.com/pingcap/tidb/pkg/testkit"
)

func TestIssue42084(t *testing.T) {
	store := testkit.CreateMockStore(t)
	tk := testkit.NewTestKit(t, store)

	tk.MustExec("use test")
	tk.MustExec("CREATE TABLE `t` ( `a` int(11) NOT NULL AUTO_INCREMENT, PRIMARY KEY (`a`) );")
	tk.MustExec("insert into t values(12);")
	tk.MustQuery("select last_insert_id();").Check(testkit.Rows("13"))
}
