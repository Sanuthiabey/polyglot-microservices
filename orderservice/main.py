from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Order(BaseModel):
    item: str
    quantity: int
    customerId: str

orders = [
    {"id": 1, "item": "Laptop", "quantity": 1, "customerId": "C001"}
]

@app.get("/orders")
def get_orders():
    return orders

@app.get("/orders/{id}")
def get_order(id: int):
    for order in orders:
        if order["id"] == id:
            return order
    return {"error": "Not found"}

@app.post("/orders")
def create_order(order: Order):
    new_order = order.dict()
    new_order["id"] = len(orders) + 1
    orders.append(new_order)
    return new_order