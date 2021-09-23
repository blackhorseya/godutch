import {call, put} from 'redux-saga/effects';
import {
  requestGetMe,
  requestLogin,
  requestLogout,
  requestSignup,
} from '../requests/user';
import {
  getMeFail,
  getMeSuccess,
  loginFail,
  loginSuccess,
  logoutFail,
  logoutSuccess,
  signupFail,
  signupSuccess,
} from '../../ducks/userSlice';

export function* handleLogin(action) {
  const {email, password} = action.payload;

  try {
    const resp = yield call(requestLogin, email, password);
    const {data} = resp;
    yield put(loginSuccess(data));
  } catch (error) {
    yield put(loginFail(error.response.data));
  }
}

export function* handleSignup(action) {
  const {email, password, name} = action.payload;

  try {
    const resp = yield call(requestSignup, email, password, name);
    const {data} = resp;
    yield put(signupSuccess(data));
  } catch (error) {
    yield put(signupFail(error.response.data));
  }
}

export function* handleGetMe(action) {
  const {token} = action.payload;

  try {
    const resp = yield call(requestGetMe, token);
    const {data} = resp;
    yield put(getMeSuccess(data));
  } catch (error) {
    yield put(getMeFail(error.response.data));
  }
}

export function* handleLogout(action) {
  const {token} = action.payload;

  try {
    const resp = yield call(requestLogout, token);
    const {data} = resp;
    yield put(logoutSuccess(data));
  } catch (error) {
    yield put(logoutFail(error.response.data));
  }
}
