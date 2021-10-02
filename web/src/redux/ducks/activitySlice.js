import {createSlice} from '@reduxjs/toolkit';

const initialState = {
  loading: false,
  activities: [],
  error: '',
};

const activitySlice = createSlice({
  name: 'activity',
  initialState: initialState,
  reducers: {
    listActivitiesRequest: (state, action) => {
      return {...state, loading: true};
    },
    listActivitiesSuccess: (state, action) => {
      return {
        ...state,
        loading: false,
        activities: action.payload.data,
        error: '',
      };
    },
    listActivitiesFail: (state, action) => {
      return {...state, loading: false, error: action.payload.msg};
    },
  },
});

export const {
  listActivitiesRequest,
  listActivitiesSuccess,
  listActivitiesFail,
} = activitySlice.actions;

export default activitySlice.reducer;
