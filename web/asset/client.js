/**
 * 虚拟用户
 */
let fictitiousUser = {
    _c: [],
    addOne: function () {
        let tmp = new WebSocket(wsAddr)
        tmp.onopen = function () {
            let that = this
            let heartbeatIndex = 0
            heartbeatIndex = setInterval(function () {
                try {
                    that.send("~3yPvmnz~")
                } catch (e) {
                    clearInterval(heartbeatIndex)
                    heartbeatIndex = 0
                }
            }, 1000 * heartbeatInterval)
        }
        fictitiousUser._c.push(tmp)
    },
    addMore: function (num) {
        if (num <= 0) {
            return
        }
        num--
        fictitiousUser.addOne()
        setTimeout(function () {
            fictitiousUser.addMore(num)
        }, 100)
    }
}
/**
 * 顶部消息
 */
let messageObj = {
    _obj: null,
    _init: function () {
        if (!messageObj._obj) {
            messageObj._obj = document.querySelector("#j-message")
        }
    },
    error: function (message) {
        messageObj._init()
        messageObj._obj.innerHTML = message
        messageObj._obj.style.color = "#ef0634"
    },
    success: function (message) {
        messageObj._init()
        messageObj._obj.innerHTML = message
        messageObj._obj.style.color = "rgb(255, 247, 126)"
    }
}
/**
 * 弹幕
 */
let barrageObj = {
    _runIndex: 0,
    _obj: null,
    _lock: false,
    _init: function () {
        if (!barrageObj._obj) {
            barrageObj._obj = document.querySelector("#j-barrage")
        }
    },
    _channelLow: [],
    _channelHigh: [],
    producer: function (message, jumping) {
        if (jumping) {
            this._channelHigh.push(message)
            return
        }
        this._channelLow.push(message)
    },
    stopConsume: function () {
        if (this._runIndex > 0) {
            clearInterval(this._runIndex)
            this._runIndex = 0
            barrageObj._obj.className = ""
            barrageObj._obj.innerHTML = ""
            barrageObj._lock = false
        }
    },
    startConsume: function () {
        if (this._runIndex > 0) {
            return
        }
        this._init()
        this._runIndex = setInterval(function () {
            if (barrageObj._lock) {
                return
            }
            let message = barrageObj._channelHigh.pop()
            if (message === undefined) {
                message = barrageObj._channelLow.pop()
            }
            if (message !== undefined) {
                barrageObj._lock = true
                barrageObj._obj.className = "barrage-animation"
                barrageObj._obj.addEventListener("animationend", () => {
                    barrageObj._obj.className = ""
                    barrageObj._obj.innerHTML = ""
                    barrageObj._lock = false
                })
                barrageObj._obj.innerHTML = message
            }
        }, 1000)
    }
}
/**
 * 3d球
 */
let thirdBrother = {
    _container: null,
    setContainer: function (container) {
        this._container = container
    },
    setRadius: function () {
        this._radius =
            parseInt((document.documentElement.clientWidth / 2).toFixed(0)) -
            parseInt(
                (
                    parseInt(
                        (document.documentElement.clientWidth / 2).toFixed(0)
                    ) / 10
                ).toFixed(0)
            )
        this._cos = Math.cos(Math.PI / this._radius)
        this._sin = Math.sin(Math.PI / this._radius)
    },
    _cos: Math.cos(Math.PI / 300),
    _sin: Math.sin(Math.PI / 300),
    _radius: 300,
    _animateIndex: 0,
    _animateFastIndex: 0,
    _tags: [],
    countTag: function () {
        let tagEle = document.querySelectorAll(".tag")
        if (tagEle) {
            return tagEle.length
        }
        return 0
    },
    clearTag: function () {
        let tagEle = document.querySelectorAll(".tag")
        if (tagEle) {
            for (let i = 0; i < tagEle.length; i++) {
                tagEle[i].remove()
            }
        }
    },
    addTag: function (id, num) {
        if (document.querySelector("#j-" + id)) {
            return
        }
        let ele = document.createElement("span")
        ele.innerText = num
        ele.id = "j-" + id
        ele.className = "tag"
        if (wsObj.currentUser && id === wsObj.currentUser.uniqId) {
            ele.style.fontSize = "10.31vw"
            ele.style.fontWeight = "700"
        } else {
            ele.style.fontSize = "7.46vw"
        }
        this._container.append(ele)
    },
    removeTag: function (id) {
        let tmp = document.querySelector("#j-" + id)
        if (tmp) {
            tmp.remove()
        }
    },
    move: function () {
        for (let i = 0; i < thirdBrother._tags.length; i++) {
            let ele = thirdBrother._tags[i]
            let y1 = ele.y * thirdBrother._cos - ele.z * thirdBrother._sin
            let z1 = ele.z * thirdBrother._cos + ele.y * thirdBrother._sin
            ele.y = y1
            ele.z = z1

            let x1 = ele.x * thirdBrother._cos - ele.z * thirdBrother._sin
            let z2 = ele.z * thirdBrother._cos + ele.x * thirdBrother._sin
            ele.x = x1
            ele.z = z2

            ele.move()
        }
    },
    animateStart: function () {
        thirdBrother.move()
        thirdBrother._animateIndex = requestAnimationFrame(
            thirdBrother.animateStart
        )
    },
    animateStop: function () {
        if (this._animateIndex > 0) {
            cancelAnimationFrame(this._animateIndex)
            this._animateIndex = 0
        }
    },
    animateFastStart: function () {
        this._cos = Math.cos(Math.PI / (this._radius / 5))
        this._sin = Math.sin(Math.PI / (this._radius / 5))
    },
    animateFastStop: function () {
        this._cos = Math.cos(Math.PI / this._radius)
        this._sin = Math.sin(Math.PI / this._radius)
    },
    reset: function () {
        let tag = function (ele, x, y, z) {
            this.ele = ele
            this.x = x
            this.y = y
            this.z = z
        }
        tag.prototype = {
            move: function () {
                let scale =
                    document.documentElement.clientWidth /
                    (document.documentElement.clientWidth - this.z)
                let alpha =
                    (this.z + thirdBrother._radius) / (2 * thirdBrother._radius)
                let left =
                    this.x +
                    thirdBrother._container.offsetWidth / 2 -
                    this.ele.offsetWidth / 2 +
                    "px"
                let top =
                    this.y +
                    thirdBrother._container.offsetHeight / 2 -
                    this.ele.offsetHeight / 2 +
                    "px"
                let transform =
                    "translate(" + left + ", " + top + ") scale(" + scale + ")"
                this.ele.style.opacity = (alpha + 0.5).toFixed(2)
                this.ele.style.zIndex = (scale * 100).toFixed(0)
                this.ele.style.transform = transform
            }
        }
        let tagEle = document.querySelectorAll(".tag")
        this._tags = []
        for (let i = 0; i < tagEle.length; i++) {
            let ele = tagEle[i]
            let k = -1 + (2 * (i + 1) - 1) / tagEle.length
            let a = Math.acos(k)
            let b = a * Math.sqrt(tagEle.length * Math.PI)
            let x = this._radius * Math.sin(a) * Math.cos(b)
            let y = this._radius * Math.sin(a) * Math.sin(b)
            let z = this._radius * Math.cos(a)
            let t = new tag(ele, x, y, z)
            this._tags.push(t)
            t.move()
        }
    }
}
/**
 * ws对象
 */
let wsObj = {
    currentUser: null,
    currentUserList: new Map(),
    onlineNum: 0,
    callback: new Map(),
    _socket: null,
    _heartbeatInterval: 0,
    showOnline: function () {
        try {
            messageObj.success(
                "我的号码：<strong>" +
                wsObj.currentUser.id +
                "</strong>，参与人数：" +
                wsObj.onlineNum
            )
        } catch (e) {
        }
    },
    _secretKey: null,
    setSecretKey: function (secretKey) {
        this._secretKey = secretKey
    },
    connect: function () {
        try {
            if (wsObj._socket && parseInt(wsObj._socket.readyState) === 1) {
                return
            }
            let url = wsAddr
            let protocols = null
            if (
                this._secretKey !== "" &&
                this._secretKey !== undefined &&
                this._secretKey !== null
            ) {
                url += "?connType=admin"
                let time = parseInt((new Date().getTime() / 1000).toFixed(0))
                let randStr = Math.random().toString(36).slice(-8)
                let md5Str = md5(time + randStr + this._secretKey)
                protocols = [md5Str, time, randStr]
            }
            //时间戳+随机值+服务器key
            let tmp = new WebSocket(url, protocols)
            if (!tmp) {
                return
            }
            tmp.onerror = wsObj.onerror
            tmp.onopen = wsObj.onopen
            tmp.onmessage = wsObj.onmessage
            tmp.onclose = wsObj.onclose
            wsObj._socket = tmp
        } catch (e) {
            console.error(e)
        }
    },
    send: function (cmd, data) {
        if (!wsObj._socket) {
            return
        }
        if (!data) {
            data = {}
        }
        //这里写死workerId
        wsObj._socket.send(
            workerId +
            JSON.stringify({
                cmd: cmd,
                data: JSON.stringify(data)
            })
        )
    },
    close: function () {
        try {
            wsObj._socket && wsObj._socket.close()
        } catch (e) {
        }
    },
    heartbeat: function () {
        try {
            if (!wsObj._socket || parseInt(wsObj._socket.readyState) !== 1) {
                return false
            }
            wsObj._socket.send("~3yPvmnz~")
            return true
        } catch (e) {
            return false
        }
    },
    onerror: function () {
        wsObj.close()
    },
    onopen: function () {
        if (wsObj._heartbeatInterval > 0) {
            clearInterval(wsObj._heartbeatInterval)
            wsObj._heartbeatInterval = 0
        }
        wsObj._heartbeatInterval = setInterval(function () {
            wsObj.heartbeat()
        }, 1000 * heartbeatInterval)
    },
    onclose: function () {
        if (wsObj._heartbeatInterval > 0) {
            clearInterval(wsObj._heartbeatInterval)
            wsObj._heartbeatInterval = 0
        }
        wsObj._socket = null
        thirdBrother.animateStop()
        thirdBrother.clearTag()
        setTimeout(function () {
            wsObj.connect()
        }, 1000)
    },
    onmessage: function (evt) {
        if (!evt.data) {
            console.error(evt.data)
            return
        }
        if (evt.data === "~u38NvZ~") {
            return
        }
        /**
         *
         * @type {{cmd: number, version: number, data:any}}
         */
        let ret = JSON.parse(evt.data)
        let cb = wsObj.callback.get(ret.cmd)
        if (cb) {
            cb(ret)
        } else {
            console.log(ret)
        }
    }
}
/**
 * 消息版本校验
 */
let versionObj = {
    _m: new Map(),
    check: function (currentCmd, version, otherCmdArr) {
        if (otherCmdArr === undefined) {
            otherCmdArr = [currentCmd]
        } else {
            otherCmdArr.push(currentCmd)
        }
        let tmp = []
        otherCmdArr.forEach(function (cmd) {
            if (versionObj._m.has(cmd)) {
                tmp.push(versionObj._m.get(cmd))
            }
        })
        tmp.push(version)
        if (Math.max(...tmp) === version) {
            versionObj._m.set(currentCmd, version)
            return true
        }
        return false
    }
}
/**
 * 与服务端约定的指令
 */
let api = {
    ConnOpen: 1,
    ConnOpenBroadcast: 2,
    ConnCloseBroadcast: 3,
    LotteryStart: 4,
    LotteryEnd: 5,
    Barrage: 6
}
/**
 * 连接打开
 */
wsObj.callback.set(api.ConnOpen, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    //记录当前的用户信息、当前的用户列表
    wsObj.currentUser = resp.data.data.user
    wsObj.currentUserList = new Map()
    wsObj.onlineNum = resp.data.data.onlineNum
    resp.data.data.userList.forEach(function (user) {
        wsObj.currentUserList.set(user.uniqId, user.id)
    })
    //渲染3d球
    thirdBrother.animateStop()
    thirdBrother.clearTag()
    wsObj.currentUserList.forEach(function (id, uniqId) {
        thirdBrother.addTag(uniqId, id)
    })
    thirdBrother.reset()
    thirdBrother.animateStart()
    //渲染在线人数
    wsObj.showOnline()
    //显示抽奖按钮
    if (wsObj.currentUser.isAdmin) {
        lotteryObj.showBtn()
    } else {
        //非管理员，显示弹幕按钮
        sendObj.showBtn()
    }
    barrageObj.producer(resp.data.data.pageWelcome, true);
})
/**
 * 其他用户连接进来
 */
wsObj.callback.set(api.ConnOpenBroadcast, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    //忽略自己的上线广播
    if (
        wsObj.currentUser &&
        resp.data.data.uniqId === wsObj.currentUser.uniqId
    ) {
        return
    }
    //在线人数自增
    wsObj.onlineNum += 1
    //如果不存在，并且当前的3d球渲染tag数比较少，则进行登记操作
    if (
        !wsObj.currentUserList.has(resp.data.data.uniqId) &&
        thirdBrother.countTag() < limit3D
    ) {
        //记录到当前用户列表
        wsObj.currentUserList.set(resp.data.data.uniqId, resp.data.data.id)
        //渲染3d球
        thirdBrother.animateStop()
        thirdBrother.addTag(resp.data.data.uniqId, resp.data.data.id)
        thirdBrother.reset()
        thirdBrother.animateStart()
    }
    //渲染在线人数
    wsObj.showOnline()
})
/**
 * 其他用户关闭
 */
wsObj.callback.set(api.ConnCloseBroadcast, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    //自减在线人数
    wsObj.onlineNum -= 1
    if (wsObj.onlineNum < 0) {
        wsObj.onlineNum = 0
    }
    //如果存在，则进行移除操作
    if (wsObj.currentUserList.has(resp.data.data.uniqId)) {
        //当前用户列表移除掉
        wsObj.currentUserList.delete(resp.data.data.uniqId)
        //渲染3d球
        thirdBrother.animateStop()
        thirdBrother.removeTag(resp.data.data.uniqId)
        thirdBrother.reset()
        thirdBrother.animateStart()
    }
    //渲染在线人数
    wsObj.showOnline()
})
/**
 * 抽奖开始
 */
wsObj.callback.set(api.LotteryStart, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    //检查服务器给出的最新的状态
    if (!versionObj.check(resp.cmd, resp.version, [api.LotteryEnd])) {
        return
    }
    //设置抽奖按钮状态
    lotteryObj.setRunning()
    //提高3d球的运转速度
    thirdBrother.animateFastStart()
})
/**
 * 抽奖关闭
 */
wsObj.callback.set(api.LotteryEnd, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    //检查服务器给出的最新的状态，因为网络原因，可能导致抽奖的状态是错乱到达的
    if (!versionObj.check(resp.cmd, resp.version, [api.LotteryStart])) {
        return
    }
    //设置抽奖按钮状态
    lotteryObj.setRunOk()
    //恢复3d球的运转速度
    thirdBrother.animateFastStop()
    //显示中奖人
    if (resp.data.data.winner.uniqId === wsObj.currentUser.uniqId) {
        //对于中奖人，直接alert，阻塞整个页面，避免网络断开后，中奖记录丢失
        if (wsObj.currentUser.isAdmin) {
            alert("恭喜获得本轮大奖！")
        } else {
            alert(
                "恭喜获得本轮大奖！\n别动，请速速兑奖！\n中奖识别码是：" +
                resp.data.data.winner.uniqId
            )
        }
    } else {
        if (wsObj.currentUser.isAdmin) {
            alert("本轮大奖识别码是：" + resp.data.data.winner.uniqId)
        } else {
            barrageObj.producer(
                "恭喜号码：" + resp.data.data.winner.id + "，获得本轮大奖！",
                true
            )
        }
    }
})
/**
 * 弹幕
 */
wsObj.callback.set(api.Barrage, function (resp) {
    if (resp.data.code !== 0) {
        messageObj.error(resp.data.message)
        return
    }
    barrageObj.producer(resp.data.data.message)
})
/**
 * 管理抽奖的对象
 */
let lotteryObj = {
    _obj: null,
    _init: function () {
        if (!lotteryObj._obj) {
            lotteryObj._obj = document
                .querySelector("#j-lottery")
                .querySelector("img")
        }
    },
    showBtn: function () {
        lotteryObj._init()
        lotteryObj._obj.parentNode.style.display = "inline-block"
    },
    setRunning: function () {
        lotteryObj._init()
        lotteryObj._obj.style.filter = "grayscale(0%)"
        lotteryObj._obj.setAttribute("data-lock", "true")
        //显示在线人数
        wsObj.showOnline()
    },
    setRunOk: function () {
        lotteryObj._init()
        lotteryObj._obj.style.filter = "grayscale(50%)"
        lotteryObj._obj.setAttribute("data-lock", "false")
    },
    _lockShake: 0,
    run: function () {
        let t = new Date().getTime()
        //锁定一会儿，让服务器将操作广播给其它用户，其它用户的状态变更后，再进行下一次操作，
        if (lotteryObj._lockShake + 1000 * 3 > t) {
            return
        }
        lotteryObj._lockShake = t
        lotteryObj._init()
        //添加抖动动画，避免用户认为没点击中
        lotteryObj._obj.className = "shake"
        //一秒后取消抖动
        setTimeout(function () {
            lotteryObj._obj.className = ""
        }, 1000)
        if (lotteryObj._obj.getAttribute("data-lock") === "true") {
            wsObj.send(api.LotteryEnd)
        } else {
            wsObj.send(api.LotteryStart)
        }
    }
}
/**
 * 发送弹幕
 */
let sendObj = {
    _obj: null,
    _init: function () {
        if (!this._obj) {
            this._obj = document.querySelector("#j-send")
        }
    },
    showBtn: function () {
        this._init()
        this._obj.style.display = "block"
    },
    _lock: 0,
    sendBarrage: function () {
        this._init()
        let input = this._obj.querySelector("input")
        if (input.value === "") {
            return
        }
        let t = new Date().getTime()
        if (t - this._lock < 5000) {
            return
        }
        this._lock = t
        if (!wsObj.currentUser) {
            return
        }
        wsObj.send(api.Barrage, {
            message: input.value,
            id: wsObj.currentUser.id
        })
        input.value = ""
        let btn = this._obj.querySelector("button")
        let index = 0
        let count = 5
        btn.innerText = count + "秒"
        index = setInterval(function () {
            if (count === 0) {
                clearInterval(index)
                btn.innerText = "发送"
                return
            }
            btn.innerText = count + "秒"
            count--
        }, 1000)
    }
}