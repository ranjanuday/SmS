import {
  useEffect,
  useState,
} from "react"

import API from "../services/api"

const useDashboard = () => {

  const [stats, setStats] = useState({})
  const [user, setUser] = useState(null)

  useEffect(() => {

    fetchStats()
    fetchUser()

  }, [])

  const fetchStats = async () => {

    const res = await API.get(
      "/dashboard/stats"
    )

    setStats(res.data)
  }

  const fetchUser = async () => {

    const token =
      localStorage.getItem("token")

    const payload = JSON.parse(
      atob(token.split(".")[1])
    )

    const res = await API.get(
      `/users/${payload.userID}`
    )

    setUser(res.data)
  }

  return {
    stats,
    user,
  }
}

export default useDashboard