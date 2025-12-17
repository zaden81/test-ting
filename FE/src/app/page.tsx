"use client";

import React, { useEffect, useRef, useState } from "react";
import "./page.css";

type Kol = {
  kolID: number;
  userProfileID: number;
  language: string;
  education: string;
  expectedSalary: number;
  code: string;
  active: boolean;
  activeDate: string;
};

type KolVM = {
  result: string;
  errorMessage: string;
  pageIndex: number;
  pageSize: number;
  guid: string;
  totalCount: number;
  kol: Kol[];
};

export default function Page() {
  const [vm, setVm] = useState<KolVM | null>(null);
  const [loading, setLoading] = useState(true);

  const listRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const run = async () => {
      setLoading(true);
      try {
        const res = await fetch("/api/kols?pageIndex=1&pageSize=30", {
          cache: "no-store",
        });
        const data: KolVM = await res.json();
        setVm(data);
      } catch {
        setVm({
          result: "Unsuccess",
          errorMessage: "Fetch failed",
          pageIndex: 1,
          pageSize: 30,
          guid: "",
          totalCount: 0,
          kol: [],
        });
      } finally {
        setLoading(false);
      }
    };
    run();
  }, []);

  const scrollUp = () => {
    listRef.current?.scrollBy({ top: -240, behavior: "smooth" });
  };

  const scrollDown = () => {
    listRef.current?.scrollBy({ top: 240, behavior: "smooth" });
  };

  return (
    <>
      <h1 className="header">KOLs List</h1>

      <div className="wrap">
        <div className="toolbar">
          <div>
            <div className="title">KOL Directory</div>
            <div className="sub">
              {loading
                ? "Loading..."
                : vm?.result === "Success"
                ? `Total: ${vm.totalCount}`
                : `Error: ${vm?.errorMessage ?? "Unknown"}`}
            </div>
          </div>

          <div className="buttons">
            <button className="btn" onClick={scrollUp} aria-label="Scroll up">
              ▲
            </button>
            <button className="btn" onClick={scrollDown} aria-label="Scroll down">
              ▼
            </button>
          </div>
        </div>

        <div className="list" ref={listRef}>
          {(vm?.kol ?? []).map((k) => (
            <div className="item" key={k.kolID}>
              <div className="itemTop">
                <div className="code">{k.code}</div>
                <div className={`badge ${k.active ? "on" : "off"}`}>
                  {k.active ? "Active" : "Inactive"}
                </div>
              </div>

              <div className="meta">
                <div>
                  <span>KolID:</span> {k.kolID}
                </div>
                <div>
                  <span>UserProfileID:</span> {k.userProfileID}
                </div>
                <div>
                  <span>Language:</span> {k.language}
                </div>
                <div>
                  <span>Education:</span> {k.education}
                </div>
                <div>
                  <span>ExpectedSalary:</span> {k.expectedSalary}
                </div>
                <div>
                  <span>ActiveDate:</span>{" "}
                  {k.activeDate ? new Date(k.activeDate).toLocaleString() : "-"}
                </div>
              </div>
            </div>
          ))}

          {!loading && (vm?.kol?.length ?? 0) === 0 && (
            <div className="empty">No KOLs found.</div>
          )}
        </div>
      </div>
    </>
  );
}
