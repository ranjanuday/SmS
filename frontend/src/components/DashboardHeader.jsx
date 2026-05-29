import React from "react"

const DashboardHeader = ({
  user,
}) => {

  return (

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
          {user?.role}
        </h2>

      </div>

    </div>
  )
}

export default DashboardHeader