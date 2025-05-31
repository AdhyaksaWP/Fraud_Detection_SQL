import warnings
warnings.filterwarnings("ignore")

import sys
import joblib
import numpy as np

model = joblib.load("../ml-model/random_forest_model.pkl")

# Get amount from command line argument
amount = float(sys.argv[1])

X = np.array([[amount]])
prediction = model.predict(X)

print(int(prediction[0]))
