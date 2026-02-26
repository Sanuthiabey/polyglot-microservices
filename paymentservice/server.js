const express = require('express');
const app = express();
app.use(express.json());

let payments = [
    { id: 1, orderId: 1, amount: 999.99, method: "CARD", status: "SUCCESS" }
];

app.get('/payments', (req, res) => {
    res.json(payments);
});

app.get('/payments/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const payment = payments.find(p => p.id === id);
    if (!payment) return res.status(404).json({ message: "Not found" });
    res.json(payment);
});

app.post('/payments/process', (req, res) => {
    const newPayment = {
        id: payments.length + 1,
        ...req.body,
        status: "SUCCESS"
    };
    payments.push(newPayment);
    res.json(newPayment);
});

app.listen(8083, () => {
    console.log("Payment service running on port 8083");
});