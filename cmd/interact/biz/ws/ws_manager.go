package ws

var WsConnManager *Manager

const (
	Matching = iota
	Free
	Gaming
)

type MetaData struct {
	Conn   *Connection
	Status int
}

type Manager struct {
	metaData    map[int64]*MetaData
	opponentMap map[int64]int64
}

func Init() {
	WsConnManager = &Manager{
		metaData:    make(map[int64]*MetaData),
		opponentMap: make(map[int64]int64),
	}
}

func (m *Manager) Put(userId int64, metaData *MetaData) {
	m.metaData[userId] = metaData
}

func (m *Manager) Get(userId int64) *MetaData {
	return m.metaData[userId]
}

func (m *Manager) Remove(userId int64) {
	delete(m.metaData, userId)
	delete(m.opponentMap, userId)
}

func (m *Manager) SetOpponent(userId, opponent int64) {
	m.opponentMap[userId] = opponent
}

func (m *Manager) SetStatus(userId int64, status int) {
	m.metaData[userId].Status = status
}

func (m *Manager) GetOpponentMetaData(userId int64) *MetaData {
	return m.Get(m.opponentMap[userId])
}
