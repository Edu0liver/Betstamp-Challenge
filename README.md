# Betstamp-Challenge

Simply running the following command on the root project:

```bash
go run .
```

It will run the function that consumes the example json file, the function assusmes that the json can have several objects in the events array, and in each object, can have several markets. That's why I choose a asynchronous approach, that each event and each market inside the events are execute in goroutines, asynchronously, and handled by workers.

Output example:

```bash
Worker id setted: 5
Worker id setted: 4
Worker id setted: 2
Worker id setted: 1
Worker id setted: 3

Number of markets processed: 6

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Total Is_live:true Odds:1.91 Number:223.5 Side_type:Over}

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Total Is_live:true Odds:1.91 Number:223.5 Side_type:Under}

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Moneyline Is_live:true Odds:2.9 Number:0 Side_type:CHA Hornets}

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Moneyline Is_live:true Odds:1.37 Number:0 Side_type:LA Lakers}

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Spread Is_live:true Odds:1.91 Number:7.5 Side_type:CHA Hornets}

{Fixture_id:CHA Hornets_%_LA Lakers_%_2025-02-20 03:00:00 +0000 UTC Bet_type:Spread Is_live:true Odds:1.91 Number:-7.5 Side_type:LA Lakers}
```
