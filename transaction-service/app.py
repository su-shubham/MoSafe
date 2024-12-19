from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/api/transactions', methods=['GET'])
def get_transactions():
    return jsonify({"message": "Transaction service is running"})

@app.route('/api/transactions/<transaction_id>', methods=['GET'])
def get_transaction(transaction_id):
    return jsonify({
        "id": transaction_id,
        "amount": 100.00,
        "type": "payment",
        "status": "completed"
    })

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8082)
