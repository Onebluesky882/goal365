"use client";
import { useState, useEffect } from "react";

const tabs = [
  {
    page: "home",
    label: "หน้าแรก",
    icon: (
      <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </svg>
    ),
  },
  {
    page: "search",
    label: "ค้นหา",
    icon: (
      <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
        <path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" />
      </svg>
    ),
  },
  {
    page: "notifications",
    label: "แจ้งเตือน",
    icon: (
      <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 22c1.1 0 2-.9 2-2h-4c0 1.1.89 2 2 2zm6-6v-5c0-3.07-1.64-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.63 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2z" />
      </svg>
    ),
    notificationCount: 3,
  },
  {
    page: "favorites",
    label: "รายการโปรด",
    icon: (
      <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
      </svg>
    ),
  },
  {
    page: "profile",
    label: "โปรไฟล์",
    icon: (
      <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
      </svg>
    ),
  },
];

export default function Footer() {
  const [activeTab, setActiveTab] = useState("home");
  const [showMenu, setShowMenu] = useState(true);
  const [lastScrollY, setLastScrollY] = useState(0);

  useEffect(() => {
    function handleScroll() {
      const currentScrollY = window.scrollY;
      if (currentScrollY > lastScrollY && currentScrollY > 100) {
        setShowMenu(false);
      } else {
        setShowMenu(true);
      }
      setLastScrollY(currentScrollY);
    }
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, [lastScrollY]);

  const handleTabClick = (tab: string) => {
    setActiveTab(tab);

    if (navigator.vibrate) {
      navigator.vibrate(50);
    }
  };

  return (
    <div className="min-h-scree bg-inherit">
      <nav
        className={`fixed bottom-0 left-0 right-0 bg-white bg-opacity-95 backdrop-blur-lg border-t border-white border-opacity-20 shadow-lg transition-transform duration-300 ${
          showMenu ? "translate-y-0" : "translate-y-full"
        }`}
      >
        <div className="max-w-md mx-auto flex justify-around items-center py-3">
          {tabs.map(({ page, label, icon, notificationCount }) => (
            <button
              key={page}
              onClick={() => handleTabClick(page)}
              className={`flex flex-col items-center text-sm font-medium rounded-xl px-3 py-2 relative transition-colors duration-300 ${
                activeTab === page
                  ? "text-indigo-600 bg-indigo-100 transform -translate-y-0.5"
                  : "text-gray-600 hover:text-indigo-600 hover:bg-indigo-50"
              }`}
              aria-current={activeTab === page ? "page" : undefined}
              type="button"
            >
              <span
                className={`w-6 h-6 mb-1 transition-transform duration-300 ${
                  activeTab === page ? "scale-110" : ""
                }`}
              >
                {icon}
              </span>
              <span>{label}</span>

              {notificationCount && notificationCount > 0 && (
                <span className="absolute top-1 right-2 bg-red-500 text-white text-xs w-5 h-5 rounded-full flex items-center justify-center animate-pulse">
                  {notificationCount}
                </span>
              )}

              {activeTab === page && (
                <span className="absolute top-0 left-1/2 transform -translate-x-1/2 w-1.5 h-1.5 bg-indigo-600 rounded-full animate-bounce"></span>
              )}
            </button>
          ))}
        </div>
      </nav>
    </div>
  );
}
