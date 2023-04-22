import pickle
import pandas as pd

def load_model():
    with open('training/trained_model.pkl', 'rb') as f:
        model = pickle.load(f)
    return model

def predict(model, input_data):
    input_df = pd.DataFrame(input_data, columns=['tld', 'dl', 'nos', 'nod', 'noh'])
    prediction = model.predict(input_df)
    return prediction[0]

if __name__ == '__main__':
    # Load the trained model
    model = load_model()

    # Example input data (preprocessed)
    input_data = [[2, 1, 0, 0], [17, 1, 0, 0]]

    # Make a prediction
    for data in input_data:
        result = predict(model, [data])
        print(f"Domain is {'malicious' if result else 'not malicious'}")
