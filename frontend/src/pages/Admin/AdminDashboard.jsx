import React from "react"

import {
  Users,
  GraduationCap,
  BookOpen,
  Bell,
  Briefcase,
  ClipboardList,
} from "lucide-react"

import Sidebar from "../../components/Sidebar"
import DashboardHeader from "../../components/DashboardHeader"
import StatCard from "../../components/StatCard"

import useDashboard from "../../hooks/useDashboard"

const AdminDashboard = () => {

  const {
    stats,
    user,
  } = useDashboard()

  return (

    <div className="min-h-screen bg-gray-950 text-white">

      <div className="flex">

        <Sidebar />

        <main className="flex-1 p-10">

          <DashboardHeader
            user={user}
          />

          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">

            <StatCard
              title="Students"
              value={stats.students}
              icon={<Users size={35} />}
              color="bg-indigo-600"
            />

            <StatCard
              title="Marks"
              value={stats.marks}
              icon={<GraduationCap size={35} />}
              color="bg-green-600"
            />

            <StatCard
              title="Assignments"
              value={stats.assignments}
              icon={<BookOpen size={35} />}
              color="bg-yellow-500"
            />

            <StatCard
              title="Notices"
              value={stats.notices}
              icon={<Bell size={35} />}
              color="bg-pink-600"
            />

            <StatCard
              title="Placements"
              value={stats.placements}
              icon={<Briefcase size={35} />}
              color="bg-cyan-600"
            />

            <StatCard
              title="Queries"
              value={stats.queries}
              icon={<ClipboardList size={35} />}
              color="bg-red-500"
            />

          </div>

        </main>

      </div>

    </div>
  )
}

export default AdminDashboard