import { NextResponse } from "next/server";

export async function GET(req: Request) {
  const { searchParams } = new URL(req.url);

  const pageIndex = searchParams.get("pageIndex") ?? "1";
  const pageSize = searchParams.get("pageSize") ?? "10";

  const beUrl = `http://localhost:8081/kols?pageIndex=${encodeURIComponent(
    pageIndex
  )}&pageSize=${encodeURIComponent(pageSize)}`;

  try {
    const res = await fetch(beUrl, { cache: "no-store" });
    const data = await res.json();
    return NextResponse.json(data, { status: res.status });
  } catch {
    return NextResponse.json(
      { result: "Unsuccess", errorMessage: "Cannot reach backend", kol: [] },
      { status: 502 }
    );
  }
}
