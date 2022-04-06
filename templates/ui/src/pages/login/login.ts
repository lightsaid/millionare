// 测试引入css是否正常显示样式
import "./login.less"
import picture from "../../assets/8921fdfb2b91ce5c29b5.jpg"

// 测试图片
const imgEle = document.createElement("img")
imgEle.src = picture
document.body.appendChild(imgEle)

// 测试 ES 代码是否背转换成 es5
function timeout(ms:any) {
    return new Promise((resolve) => {
      setTimeout(resolve, ms);
    });
  }
  
  async function asyncPrint(value:any, ms:any) {
    await timeout(ms);
    console.log(value);
  }
  
  asyncPrint('hello world ！', 50);