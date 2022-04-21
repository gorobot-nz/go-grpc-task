# GRPC-TASK

## pkg/server
challengeServiceServer implements UnimplementedChallengeServiceServer and has two side-APIs client for Bilty and Timer

Server has ReadMetadata, MakeShortLink, StartTimer methods

P.S. I failed Bitly interaction by API, so I decided to use side Bitly client. Timer response return 0 remaining seconds. I tried different types to get remaining_second from json (int, *int, json.Number)

## pkg/client
Client has shortLink, readMetadata, startTimer, help methods. If user enter wrong command, help method returns commands templates.
