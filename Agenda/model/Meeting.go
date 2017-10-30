package model

type Meeting struct {
    Sponsor string
    Participators []string
    StartDate string
    EndDate string
    Title string
}

func (m *Meeting) meetingInit(t_sponsor string, t_participators []string, t_startDate string, t_endDate string, t_title string) {
    m.Sponsor = t_sponsor
    m.Participators = t_participators
    m.StartDate = t_startDate
    m.EndDate = t_endDate
    m.Title = t_title
}

func (m Meeting) getSponsor() string {
    return m.Sponsor
}

func (m *Meeting) setSponsor(t_sponsor string) {
    m.Sponsor = t_sponsor
}

func (m Meeting) getParticipators() []string {
    return m.Participators
}

func (m *Meeting) setParticipators(t_participators []string) {
    m.Participators = t_participators
}

func (m Meeting) getStartDate() string {
    return m.StartDate
}

func (m *Meeting) setStartDate(t_startDate string) {
    m.StartDate = t_startDate
}

func (m Meeting) getEndDate() string {
    return m.EndDate
}

func (m *Meeting) setEndDate(t_endDate string) {
    m.EndDate = t_endDate
}

func (m Meeting) getTitle(t_title string) {
    m.Title = t_title
}

func (m *Meeting) IsParticipator(t_name string) bool {
    for _, find_p := range m.Participators {
        if find_p == t_name {
            return true
        }
    }
    return false
}
