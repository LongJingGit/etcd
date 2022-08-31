// Copyright 2019 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracker

// StateType is the state of a tracked follower.
type StateType uint64

// leader 节点保存的 follower 节点的状态。关于这些状态的解释见: 《codedump 的网络日志: etcd Raft 库解析》
const (
	// StateProbe indicates a follower whose last index isn't known. Such a
	// follower is "probed" (i.e. an append sent periodically) to narrow down
	// its last index. In the ideal (and common) case, only one round of probing
	// is necessary as the follower will react with a hint. Followers that are
	// probed over extended periods of time are often offline.
	StateProbe StateType = iota // 探测状态
	// StateReplicate is the state steady in which a follower eagerly receives
	// log entries to append to its log.
	StateReplicate // 接收日志同步请求的正常状态
	// StateSnapshot indicates a follower that needs log entries not available
	// from the leader's Raft log. Such a follower needs a full snapshot to
	// return to StateReplicate.
	// 节点在同步快照数据的状态. 这时候说明该节点已经落后 leader 数据比较多才采用了接收快照数据的状态。
	// 在节点落后 leader 数据很多的情况下，可能 leader 会多次通过 snapshot 同步数据给节点，而当 pr.Match >= pr.PendingSnapshot 的时候，
	// 说明通过快照来同步数据的流程完成了，这时可以进入正常的接收同步数据状态了，这就是函数 Progress.needSnapshotAbort 要做的判断。
	StateSnapshot
)

var prstmap = [...]string{
	"StateProbe",
	"StateReplicate",
	"StateSnapshot",
}

func (st StateType) String() string { return prstmap[uint64(st)] }
