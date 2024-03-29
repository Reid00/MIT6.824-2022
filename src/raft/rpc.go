package raft

import "fmt"

type RequestVoteRequest struct {
	Term         int
	CandidateId  int
	LastLogIndex int
	LastLogTerm  int
}

func (req RequestVoteRequest) String() string {
	return fmt.Sprintf("{Term: %v, CandidateId: %v, LastLogIndex: %v, LastLogTerm: %v}",
		req.Term, req.CandidateId, req.LastLogIndex, req.LastLogTerm)
}

type RequestVoteResponse struct {
	Term        int
	VoteGranted bool
}

func (resp RequestVoteResponse) String() string {
	return fmt.Sprintf("{Term: %v, VoteGranted: %v}", resp.Term, resp.VoteGranted)
}

// ------------------------------------------------

type AppendEntriesReq struct {
	Term          int
	LeaderId      int
	PrevLogIndex  int
	PrevLogTerm   int
	LeaderComment int
	Entries       []Entry
}

func (req AppendEntriesReq) String() string {
	return fmt.Sprintf("{Term: %d, LeaderId: %v, PreVoteLogIndex: %v, PreVoteLogTerm: %v, LeaderComment: %v, Entries: %v}",
		req.Term, req.LeaderId, req.PrevLogIndex, req.PrevLogTerm, req.LeaderComment, req.Entries)
}

type AppendEntriesResp struct {
	Term    int
	Success bool
	// for fast backup https://mit-public-courses-cn-translatio.gitbook.io/mit6-824/lecture-07-raft2/7.3-hui-fu-jia-su-backup-acceleration
	ConflictIndex int
	ConflictTerm  int
	ConflictLen   int
}

func (resp AppendEntriesResp) String() string {
	return fmt.Sprintf("{Term:%v,Success:%v,ConflictIndex:%v,ConflictTerm:%v}",
		resp.Term, resp.Success, resp.ConflictIndex, resp.ConflictTerm)
}

// ---------------------------------------------------

type InstallSnapshotReq struct {
	Term              int
	LeaderId          int
	LastIncludedIndex int
	LastIncludedTerm  int
	Data              []byte
	//Offset           int //此次传输chunk在快照文件的偏移量，快照文件可能很大，因此需要分chunk，此次不分片
	//Done             bool //是否最后一块
}

func (req InstallSnapshotReq) String() string {
	return fmt.Sprintf("{Term:%v,LeaderId:%v,LastIncludedIndex:%v,LastIncludedTerm:%v,DataSize:%v}",
		req.Term, req.LeaderId, req.LastIncludedIndex, req.LastIncludedTerm, len(req.Data))
}

type InstallSnapshotResp struct {
	Term int
}

func (resp InstallSnapshotResp) String() string {
	return fmt.Sprintf("{Term: %v}", resp.Term)
}
