const express = require('express'), bodyParser = require('body-parser');
const app = express();
const port = process.env.port || 3333;

app.use(bodyParser.json())


function productInfo(req, res) {
    return res.json(getProductInfo(req.params.id));
}

function getProductInfo(productId) {
    const products = [{
        id: 1, name: 'O seu Notebook foi separado'
    }, {
        id: 2, name: 'O seu Ipad foi separado'
    }, {
        id: 3, name: 'O seu Smartphone foi separado'
    }];

    const product = products.filter(p => {
        return p.id == productId
    });

    if(product.length){
        return product[0].name
    }else{
        return "Produto não identificado";
    }
}

app.get('/:id', (request, response) => {
    productInfo(request, response);
});

app.listen(port, () => {
    console.log(`server is runing on port ${port}`);
})
