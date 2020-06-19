var Web3 = require('web3');
var net = require('net');
var request = require('request');

//收账地址
const receiveAddress = "0xc7e05f9f72084158db2ec35ecbefc95ab095934d";

//设置服务器
var web3 = new Web3('/home/steven/.ethereum/geth.ipc', net);

//订阅pending状态的订单，过滤to为自己的订单
var subscription = web3.eth.subscribe('pendingTransactions', function(error, result){
    if (!error) {
        console.log('----获取转账id----');
	console.log(result);
	console.log('----获取to地址相关的id----');
	var tr = web3.eth.getTransaction(result).then(resp=>{
              if (resp.to.toLowerCase() == receiveAddress) {
	        console.log('#############');
                console.log('hash:', resp.hash);
                console.log('from:', resp.from.toLowerCase());
                console.log('to:', resp.to.toLowerCase());
                console.log('value:', resp.value);
		console.log('#############');
		DoPost(resp.hash, resp.from.toLowerCase(), resp.to.toLowerCase(), resp.value);
	     }
	  });
    }
});


// 发起http请求
function DoPost(trId, from, to, amount) {
    var options = {
	    url: "http://localhost:8080/transaction",
        qs: {
          id: 20200619
        },
        headers: {
          trId: trId
        },
        form: {
          trId: trId,
          from: from,
          to: to,
          amount: amount
        }
    }
    
    request.post(options, function(error, response, body){
        console.info('response:' + JSON.stringify(response));
        console.info('statusCode:' + response.statusCode);
        console.info('body:' + body);
    });
}
