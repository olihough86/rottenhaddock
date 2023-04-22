import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import classification_report, confusion_matrix
import pickle

# Load the preprocessed data
data = pd.read_csv('training/data/shuffled_domains.csv')

# Split the data into features (X) and labels (y)
X = data[["dl", "nos", "nod", "noh"]]
y = data['m']

# Split the data into training and testing sets
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Create and train the Random Forest Classifier
clf = RandomForestClassifier(n_estimators=100, random_state=42)
clf.fit(X_train, y_train)

# Test the model's performance
y_pred = clf.predict(X_test)
print(classification_report(y_test, y_pred))
print(confusion_matrix(y_test, y_pred))

# Save the trained model to a file
with open('training/trained_model.pkl', 'wb') as f:
    pickle.dump(clf, f)
