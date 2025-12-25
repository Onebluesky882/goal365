import type { PredictionResponse } from "@/types/fixture-analytis";

export const MOCK_DATA: PredictionResponse = {
  Items: [
    {
      FixtureID: 1486269,
      Fixture: {
        fixture: {
          id: 1486269,
          date: "2025-12-02T19:00:00+00:00",
          venue: { name: "Cardiff City Stadium", city: "Cardiff" },
          status: { long: "Match Finished", short: "FT" },
        },
        league: {
          id: 46,
          name: "EFL Trophy",
          country: "England",
          logo: "https://media.api-sports.io/football/leagues/46.png",
          flag: "https://media.api-sports.io/flags/gb-eng.svg",
          season: 2025,
        },
        teams: {
          home: {
            id: 43,
            name: "Cardiff",
            logo: "https://media.api-sports.io/football/teams/43.png",
            winner: false,
          },
          away: {
            id: 1333,
            name: "AFC Wimbledon",
            logo: "https://media.api-sports.io/football/teams/1333.png",
            winner: true,
          },
        },
        goals: { home: 1, away: 5 },
      },
      Predictions: {
        predictions: {
          winner: { id: 43, name: "Cardiff", comment: "Win or draw" },
          advice: "Double chance : Cardiff or draw",
          goals: { home: "-1.5", away: "-2.5" },
          percent: { home: "45%", draw: "45%", away: "10%" },
        },
        teams: {
          home: {
            id: 43,
            name: "Cardiff",
            logo: "https://media.api-sports.io/football/teams/43.png",
            last_5: {
              played: 5,
              form: "60%",
              att: "25%",
              def: "88%",
              goals: {
                for: { total: 7, average: "1.4" },
                against: { total: 4, average: "0.8" },
              },
            },
            league: {
              form: "WWLDW",
              goals: {
                for: { minute: { "76-90": { total: 1, percentage: "25%" } } },
                against: { minute: {} },
              },
              clean_sheet: { total: 2 },
              biggest: { wins: { home: "3-1", away: "0-1" } },
            },
          },
          away: {
            id: 1333,
            name: "AFC Wimbledon",
            logo: "https://media.api-sports.io/football/teams/1333.png",
            last_5: {
              played: 5,
              form: "80%",
              att: "38%",
              def: "56%",
              goals: {
                for: { total: 10, average: "2.0" },
                against: { total: 7, average: "1.4" },
              },
            },
            league: {
              form: "WLWWW",
              goals: {
                for: { minute: { "31-45": { total: 2, percentage: "33%" } } },
                against: { minute: {} },
              },
              clean_sheet: { total: 1 },
              biggest: { wins: { home: "3-1", away: "1-2" } },
            },
          },
        },
        comparison: {
          form: { home: "50%", away: "50%" },
          att: { home: "40%", away: "60%" },
          def: { home: "78%", away: "22%" },
          poisson_distribution: { home: "34%", away: "66%" },
          h2h: { home: "100%", away: "0%" },
          goals: { home: "100%", away: "0%" },
          total: { home: "50.5%", away: "49.5%" },
        },
        h2h: [
          {
            fixture: {
              id: 1387150,
              date: "2025-08-19T18:45:00+00:00",
              status: { short: "FT" },
              venue: { city: "London" },
            },
            teams: {
              home: {
                name: "AFC Wimbledon",
                logo: "https://media.api-sports.io/football/teams/1333.png",
              },
              away: {
                name: "Cardiff",
                logo: "https://media.api-sports.io/football/teams/43.png",
              },
            },
            goals: { home: 0, away: 1 },
          },
          {
            fixture: {
              id: 455310,
              date: "2015-08-11T18:45:00+00:00",
              status: { short: "FT" },
              venue: { city: "Caerdydd" },
            },
            teams: {
              home: {
                name: "Cardiff",
                logo: "https://media.api-sports.io/football/teams/43.png",
              },
              away: {
                name: "AFC Wimbledon",
                logo: "https://media.api-sports.io/football/teams/1333.png",
              },
            },
            goals: { home: 1, away: 0 },
          },
          {
            fixture: {
              id: 3001,
              date: "2014-07-26T14:00:00+00:00",
              status: { short: "FT" },
              venue: { city: "London" },
            },
            teams: {
              home: {
                name: "AFC Wimbledon",
                logo: "https://media.api-sports.io/football/teams/1333.png",
              },
              away: {
                name: "Cardiff",
                logo: "https://media.api-sports.io/football/teams/43.png",
              },
            },
            goals: { home: 2, away: 3 },
          },
          {
            fixture: {
              id: 2002,
              date: "2010-08-10T18:45:00+00:00",
              status: { short: "FT" },
              venue: { city: "Caerdydd" },
            },
            teams: {
              home: {
                name: "Cardiff",
                logo: "https://media.api-sports.io/football/teams/43.png",
              },
              away: {
                name: "AFC Wimbledon",
                logo: "https://media.api-sports.io/football/teams/1333.png",
              },
            },
            goals: { home: 4, away: 1 },
          },
          {
            fixture: {
              id: 1003,
              date: "2009-07-22T19:00:00+00:00",
              status: { short: "FT" },
              venue: { city: "London" },
            },
            teams: {
              home: {
                name: "AFC Wimbledon",
                logo: "https://media.api-sports.io/football/teams/1333.png",
              },
              away: {
                name: "Cardiff",
                logo: "https://media.api-sports.io/football/teams/43.png",
              },
            },
            goals: { home: 0, away: 0 },
          },
        ],
      },
    },
  ],
};
