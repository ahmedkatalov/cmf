import { CircularProgress } from "@mui/material"

const Loading = () => {
    return(
        <div className="w-screen h-screen flex justify-center items-center">
            <CircularProgress />
        </div>
    )
}

export default Loading