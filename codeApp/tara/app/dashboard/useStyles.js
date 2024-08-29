import { makeStyles } from "@mui/styles";

const useStyles = makeStyles((theme) => ({
  centerBoxText: {
    textAlign: "center",
    color: "#f3def5",
    fontWeight: "bolder",
    fontSize: 18,
    paddingLeft: 8,
  },
  titleText: {
    color: "#f3def5",
    fontWeight: "bolder",
  },
}));

export default useStyles;

export const searchTextFieldInputPropsStyle = {
  color: "#f3def5",
  fontWeight: "bolder",
  fontSize: 16,
};

export const listContainBoxStyle = {
  overflowX: "auto",
  display: "flex",
  padding: 2,
};
export const listBoxStyle = {
  height: 168,
  width: 232,
  backgroundColor: "primary.main",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  borderRadius: 2,
};

export const employerByDepartmentBoxStyle = {
  borderRadius: 4,
};

export const departmentsListBoxStyle={
  height:56,
  width:'90%',
  backgroundColor:"primary.main",
  display:"flex",
  
}