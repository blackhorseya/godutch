import {createSlice} from '@reduxjs/toolkit';

const initialState = {
  loading: false,
  profile: undefined,
  error: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState: initialState,
  reducers: {
    loginRequest: (state, action) => {
      return {...state, loading: true};
    },
    loginSuccess: (state, action) => {
      localStorage.setItem('token', action.payload.data.token);
      return {
        ...state,
        loading: false,
        profile: action.payload.data,
        error: '',
      };
    },
    loginFail: (state, action) => {
      return {...state, loading: false, error: action.payload.msg};
    },

    signupRequest: (state, action) => {
      return {...state, loading: true};
    },
    signupSuccess: (state, action) => {
      return {...state, loading: false, error: ''};
    },
    signupFail: (state, action) => {
      return {...state, loading: false, error: action.payload.msg};
    },

    getMeRequest: (state, action) => {
      return {...state, loading: true};
    },
    getMeSuccess: (state, action) => {
      return {
        ...state,
        loading: false,
        profile: action.payload.data,
        error: '',
      };
    },
    getMeFail: (state, action) => {
      return {...state, loading: false, error: action.payload.msg};
    },

    logoutRequest: (state, action) => {
      return {...state, loading: true};
    },
    logoutSuccess: (state, action) => {
      localStorage.removeItem('token');
      return {...state, loading: false, profile: undefined, error: ''};
    },
    logoutFail: (state, action) => {
      localStorage.removeItem('token');
      return {...state, loading: false, error: action.payload.msg};
    },
  },
});

export const {
  loginRequest,
  loginSuccess,
  loginFail,
  signupRequest,
  signupSuccess,
  signupFail,
  getMeRequest,
  getMeSuccess,
  getMeFail,
  logoutRequest,
  logoutSuccess,
  logoutFail,
} = userSlice.actions;

export default userSlice.reducer;