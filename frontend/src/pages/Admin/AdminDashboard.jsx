import React, {
  useEffect,
  useState,
} from "react"

import {
  Users,
  GraduationCap,
  BookOpen,
  Bell,
  Briefcase,
 ClipboardList,
  LogOut,
} from "lucide-react"

import API from "../../services/api"

const AdminDashboard = () => {

  // =========================
  // STATE
  // =========================

  const [stats, setStats] = useState({
    students: 0,
    marks: 0,
    assignments: 0,
    notices: 0,
    placements: 0,
    queries: 0,
  })
  const [user, setUser] = useState(null)

  // =========================
  // FETCH STATS
  // =========================
useEffect(() => {

  fetchDashboardStats()

  fetchUser()

}, [])

  const fetchDashboardStats = async () => {

    try {

      const response = await API.get(
        "/dashboard/stats"
      )

      console.log(response.data)

      setStats(response.data)

    } catch (error) {

      console.log(error)
    }
  }
  const fetchUser = async () => {

  try {

    const token = localStorage.getItem(
      "token"
    )

    const payload = JSON.parse(
      atob(token.split(".")[1])
    )

    const userID = payload.userID

    const response = await API.get(
      `/users/${userID}`
    )

    console.log(response.data)

    setUser(response.data)

  } catch (error) {

    console.log(error)
  }
}

  return (

    <div className="min-h-screen bg-gray-950 text-white">

      {/* ========================= */}
      {/* SIDEBAR */}
      {/* ========================= */}

      <div className="flex">

        <aside className="w-72 min-h-screen bg-gradient-to-b from-indigo-900 to-purple-900 shadow-2xl p-6 relative">

          {/* LOGO */}

          <div className="mb-12">

            <h1 className="text-3xl font-bold tracking-wide">
              SMS Admin
            </h1>

            <p className="text-indigo-200 mt-2 text-sm">
              Student Management System
            </p>

          </div>

          {/* MENU */}

          <nav className="space-y-4">

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <Users size={22} />

              <span className="text-lg">
                Students
              </span>

            </button>

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <GraduationCap size={22} />

              <span className="text-lg">
                Marks
              </span>

            </button>

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <BookOpen size={22} />

              <span className="text-lg">
                Assignments
              </span>

            </button>

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <Bell size={22} />

              <span className="text-lg">
                Notices
              </span>

            </button>

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <Briefcase size={22} />

              <span className="text-lg">
                Placements
              </span>

            </button>

            <button className="w-full flex items-center gap-4 bg-white/10 hover:bg-white/20 transition p-4 rounded-2xl">

              <ClipboardList size={22} />

              <span className="text-lg">
                Queries
              </span>

            </button>

          </nav>

          {/* LOGOUT */}

          <div className="absolute bottom-8 left-6 right-6">

            <button
              className="w-[224px] flex items-center justify-center gap-3 bg-red-600 hover:bg-red-700 transition py-4 rounded-2xl font-semibold"
              onClick={() => {

                localStorage.clear()

                window.location.href = "/"
              }}
            >

              <LogOut size={20} />

              Logout

            </button>

          </div>

        </aside>

        {/* ========================= */}
        {/* MAIN CONTENT */}
        {/* ========================= */}

        <main className="flex-1 p-10 bg-gradient-to-br from-gray-950 via-gray-900 to-black min-h-screen">

          {/* HEADER */}

          <div className="flex items-center justify-between mb-10">

            <div>

              <h1 className="text-4xl font-bold">
                Admin Dashboard
              </h1>

              <p className="text-gray-400 mt-2">
                Welcome back, {user?.name}
              </p>

            </div>

            <div className="bg-white/10 backdrop-blur-lg px-6 py-3 rounded-2xl border border-white/10">

              <p className="text-sm text-gray-300">
                Logged in as
              </p>

              <h2 className="font-semibold text-lg">
                Administrator
              </h2>

            </div>

          </div>

          {/* ========================= */}
          {/* STATS CARDS */}
          {/* ========================= */}

          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">

            {/* STUDENTS */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Total Students
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.students}
                  </h2>

                </div>

                <div className="bg-indigo-600 p-4 rounded-2xl">

                  <Users size={35} />

                </div>

              </div>

            </div>

            {/* MARKS */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Marks Uploaded
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.marks}
                  </h2>

                </div>

                <div className="bg-green-600 p-4 rounded-2xl">

                  <GraduationCap size={35} />

                </div>

              </div>

            </div>

            {/* ASSIGNMENTS */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Assignments
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.assignments}
                  </h2>

                </div>

                <div className="bg-yellow-500 p-4 rounded-2xl">

                  <BookOpen size={35} />

                </div>

              </div>

            </div>

            {/* NOTICES */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Active Notices
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.notices}
                  </h2>

                </div>

                <div className="bg-pink-600 p-4 rounded-2xl">

                  <Bell size={35} />

                </div>

              </div>

            </div>

            {/* PLACEMENTS */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Placements
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.placements}
                  </h2>

                </div>

                <div className="bg-cyan-600 p-4 rounded-2xl">

                  <Briefcase size={35} />

                </div>

              </div>

            </div>

            {/* QUERIES */}

            <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

              <div className="flex items-center justify-between">

                <div>

                  <p className="text-gray-400 text-lg">
                    Student Queries
                  </p>

                  <h2 className="text-5xl font-bold mt-4">
                    {stats.queries}
                  </h2>

                </div>

                <div className="bg-red-500 p-4 rounded-2xl">

                  <ClipboardList size={35} />

                </div>

              </div>

            </div>

          </div>

        </main>

      </div>

    </div>
  )
}

export default AdminDashboard