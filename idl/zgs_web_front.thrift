namespace go zgs_web_front

service ZgsWebFront {
    string HeartBeat() (api.get="/status")
}
