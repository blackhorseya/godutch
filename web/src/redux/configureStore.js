import {
  combineReducers,
  configureStore,
  getDefaultMiddleware,
} from '@reduxjs/toolkit';
import userReducer, {logoutSuccess} from './ducks/userSlice';
import createSagaMiddleware from 'redux-saga';
import {watcherSaga} from './sagas/rootSaga';

const sagaMiddleware = createSagaMiddleware();

const appReducer = combineReducers({
  user: userReducer,
});

const rootReducer = (state, action) => {
  if (action.type === logoutSuccess.type) {
    state = undefined;
  }

  return appReducer(state, action);
};

const store = configureStore({
  reducer: rootReducer,
  middleware: [...getDefaultMiddleware({thunk: false}), sagaMiddleware],
});

sagaMiddleware.run(watcherSaga);

export default store;
