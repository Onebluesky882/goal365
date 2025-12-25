package fixtures

import m "mytipster/models/fixture"

func getFixtureIds(fixture *m.RootFixtureResponse) ([]int, error) {
	var result = make([]int, 0, len(fixture.Response))
	for _, id := range fixture.Response {
		result = append(result, id.Fixture.ID)
	}
	return result, nil
}
