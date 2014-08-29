package rados_test

import "testing"
//import "bytes"
import "github.com/noahdesu/rados"
import "github.com/stretchr/testify/assert"
import "fmt"
import "os"

func TestVersion(t *testing.T) {
    var major, minor, patch = rados.Version()
    assert.False(t, major < 0 || major > 1000, "invalid major")
    assert.False(t, minor < 0 || minor > 1000, "invalid minor")
    assert.False(t, patch < 0 || patch > 1000, "invalid patch")
}

func TestGetSetConfigOption(t *testing.T) {
    conn, _ := rados.NewConn()

    // rejects invalid options
    err := conn.SetConfigOption("wefoijweojfiw", "welfkwjelkfj")
    assert.Error(t, err, "Invalid option")

    // verify SetConfigOption changes a values
    log_file_val, err := conn.GetConfigOption("log_file")
    assert.NotEqual(t, log_file_val, "/dev/null")

    err = conn.SetConfigOption("log_file", "/dev/null")
    assert.NoError(t, err, "Invalid option")

    log_file_val, err = conn.GetConfigOption("log_file")
    assert.Equal(t, log_file_val, "/dev/null")
}

func TestParseDefaultConfigEnv(t *testing.T) {
    conn, _ := rados.NewConn()

    log_file_val, _ := conn.GetConfigOption("log_file")
    assert.NotEqual(t, log_file_val, "/dev/null")

    err := os.Setenv("CEPH_ARGS", "--log-file /dev/null")
    assert.NoError(t, err)

    err = conn.ParseDefaultConfigEnv()
    assert.NoError(t, err)

    log_file_val, _ = conn.GetConfigOption("log_file")
    assert.Equal(t, log_file_val, "/dev/null")
}

func TestParseCmdLineArgs(t *testing.T) {
    conn, _ := rados.NewConn()

    log_file_val, _ := conn.GetConfigOption("mon_host")
    assert.NotEqual(t, log_file_val, "127.0.0.1")

    args := []string{ "--mon-host 127.0.0.1" }
    err := conn.ParseCmdLineArgs(args)
    assert.NoError(t, err)

    log_file_val, _ = conn.GetConfigOption("mon_host")
    assert.Equal(t, log_file_val, "127.0.0.1")
}

func TestGetClusterStats(t *testing.T) {
    conn, _ := rados.NewConn()
    conn.ReadDefaultConfigFile()
    conn.Connect()

    stat, err := conn.GetClusterStats()
    assert.NoError(t, err)
    assert.True(t, stat.Kb > 0)
    assert.True(t, stat.Kb_used > 0)
    assert.True(t, stat.Kb_avail > 0)
    assert.True(t, stat.Num_objects > 0)
}

func TestGetFSID(t *testing.T) {
    conn, _ := rados.NewConn()
    conn.ReadDefaultConfigFile()
    conn.Connect()

    fsid, err := conn.GetFSID()
    assert.NoError(t, err)
    assert.NotEqual(t, fsid, "")
}

func TestGetInstanceID(t *testing.T) {
    conn, _ := rados.NewConn()
    conn.ReadDefaultConfigFile()
    conn.Connect()

    id := conn.GetInstanceID()
    assert.NotEqual(t, id, 0)
}




//func TestOpen(t *testing.T) {
//    _, err := rados.NewConn()
//    assert.Equal(t, err, nil, "error")
//}
//
//func TestConnect(t *testing.T) {
//    conn, _ := rados.NewConn()
//    conn.ReadDefaultConfigFile()
//    err := conn.Connect()
//    assert.Equal(t, err, nil)
//}
//
//func TestPingMonitor(t *testing.T) {
//    conn, _ := rados.NewConn()
//    conn.ReadDefaultConfigFile()
//    conn.Connect()
//    reply, err := conn.PingMonitor("kyoto")
//    assert.Equal(t, err, nil)
//    assert.True(t, len(reply) > 0)
//}

func TestListPools(t *testing.T) {
    conn, _ := rados.NewConn()
    conn.ReadDefaultConfigFile()
    conn.Connect()
    pools, _ := conn.ListPools()
    fmt.Println(len(pools), pools)
}

func TestSetConfigOption(t *testing.T) {
    conn, _ := rados.NewConn()
    conn.ReadDefaultConfigFile()
    conn.Connect()

    err := conn.WaitForLatestOSDMap()
    assert.NoError(t, err)

    stat, err := conn.GetClusterStats()
    assert.NoError(t, err)
    assert.True(t, stat.Kb > 0)
    assert.True(t, stat.Kb_used > 0)
    assert.True(t, stat.Kb_avail > 0)
    assert.True(t, stat.Num_objects > 0)

    args := []string{ "--mon-host 127.0.0.1" }
    conn2, _ := rados.NewConn()
    err = conn2.ParseCmdLineArgs(args)
    assert.NoError(t, err)

    args = []string{ "--mmm-host 127.0.0.1" }
    err = conn2.ParseCmdLineArgs(args)
    assert.NoError(t, err)
}

//func TestConnect(t *testing.T) {
//    conn, _ := rados.Open("admin")
//    conn.ReadConfigFile("/home/nwatkins/ceph/ceph/src/ceph.conf")
//    conn.Connect()
//    pool, _ := conn.OpenPool("data")
//
//    data_in := []byte("blah");
//    data_out := make([]byte, 10)
//
//    pool.Write("xyz", data_in, 0)
//    pool.Read("xyz", data_out[:4], 0)
//
//    if !bytes.Equal(data_in, data_out[:4]) {
//        t.Errorf("yuk")
//    }
//}
