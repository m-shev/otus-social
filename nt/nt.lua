wrk.method = "GET"
wrk.path = "/user/list?name=%D0%BC%D0%B8%D1%85&surname=%D0%A8%D0%B5%D0%B2&skip=0&take=50"
--wrk.path = "/user/list?name=%D1%80%D0%BE%D0%BC&surname=%D0%B1%D0%BE%D0%B1%D1%80&skip=0&take=15"
--wrk.path = "/user/list?name=%D0%BC%D0%B8%D1%85&skip=0&take=50"

--t = {}
--init = function()
--    local temp = {
--        "/user/list?name=%D0%BC%D0%B8%D1%85&skip=0&take=50",
--        "/user/list?name=%D0%BC%D0%B8%D1%85&surname=%D1%88%D0%B5%D0%B2&skip=0&take=50"
--    }
--
--    t[1] = wrk.format(nil, temp[1])
--    t[2] = wrk.format(nil, temp[2])
--
--    i = 1
--end
--
--request = function()
--    i = i + 1
--
--    if i > 2 then
--        i = 1
--    end
--
--    return t[i]
--end