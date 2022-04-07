
import "./login.less"
import FromData from "../../common/fomrs"
import Toast from "../../common/toast"
const container = document.querySelector(".container")!,
    forget = document.querySelector("#forget")!, // 忘记密码
    register = document.querySelector(".register-link")!, // 马上注册
    login = document.querySelector(".login-link")!, // 去登录
    
    // 登录 from 下 icon密码显示隐藏
    loginPwdHidden = document.querySelector(".login-pwd-hidden")!,
    loginPwdShow = document.querySelector(".login-pwd-show")!,
    loginPassword: HTMLInputElement = document.querySelector(".login input[type=password]")!,

    // 注册下密码1
    regPwdHidden = document.querySelector(".reg-pwd-hidden1")!,
    regPwdShow = document.querySelector(".reg-pwd-show1")!,
    regPassword: HTMLInputElement = document.querySelector(".register .password1")!,

    // 确认
    reg2PwdHidden = document.querySelector(".reg-pwd-hidden2")!,
    reg2PwdShow = document.querySelector(".reg-pwd-show2")!,
    reg2Password: HTMLInputElement = document.querySelector(".register .password2")!,

    loginSubmit = document.querySelector("#login-btn")!,
    registerSubmit = document.querySelector("#register-btn")!;

register.addEventListener("click", () => {
    container.classList.add("active")
})

login.addEventListener("click", () => {
    container.classList.remove("active")
})

// 控制密码显示或者隐藏
const showHiddenPwd = (target:Element, hidden:Element, input:HTMLInputElement) => {
    target.addEventListener("click", ()=>{
        target.classList.toggle("hidden")
        hidden.classList.toggle("hidden")
        if (input.type == "password"){
            input.type = "text"
        }else{
            input.type = "password"
        }
    })
}

// 给每个密码icon添加事件
showHiddenPwd(loginPwdHidden, loginPwdShow, loginPassword)
showHiddenPwd(loginPwdShow, loginPwdHidden, loginPassword)


showHiddenPwd(regPwdHidden, regPwdShow, regPassword)
showHiddenPwd(regPwdShow, regPwdHidden, regPassword)

showHiddenPwd(reg2PwdHidden, reg2PwdShow, reg2Password)
showHiddenPwd(reg2PwdShow, reg2PwdHidden, reg2Password)


forget.addEventListener("click", ()=>{
    new Toast("功能实现中开发ing。。。", 1200)
})


loginSubmit.addEventListener("click", ()=>{
    let form: HTMLFormElement = document.querySelector(".login")!
    let values = FromData(form)
    console.log(values)
})
registerSubmit.addEventListener("click", ()=>{
    let form: HTMLFormElement = document.querySelector(".register")!
    let values = FromData(form)
    console.log(values)
})