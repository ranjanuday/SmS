import React from "react"
import {
  Users,
  GraduationCap,
  BookOpen,
  Bell,
  Briefcase,
  ClipboardList,
  LogOut,
} from "lucide-react"

const Sidebar = () => {

  const logout = () => {

    localStorage.clear()

    window.location.href = "/"
  }

  return (

    <aside className="w-72 min-h-screen bg-gradient-to-b from-indigo-900 to-purple-900 shadow-2xl p-6 relative">

      <div className="mb-12">

        <h1 className="text-3xl font-bold">
          SMS Admin
        </h1>

      </div>

      <nav className="space-y-4">

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <Users />
          Students
        </button>

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <GraduationCap />
          Marks
        </button>

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <BookOpen />
          Assignments
        </button>

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <Bell />
          Notices
        </button>

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <Briefcase />
          Placements
        </button>

        <button className="w-full flex items-center gap-4 p-4 rounded-2xl bg-white/10">
          <ClipboardList />
          Queries
        </button>

      </nav>

      <button
        onClick={logout}
        className="absolute bottom-8 left-6 right-6 bg-red-600 py-4 rounded-2xl flex justify-center gap-3"
      >
        <LogOut />
        Logout
      </button>

    </aside>
  )
}

export default Sidebar