import * as React from "react"
import { StaticImage } from "gatsby-plugin-image"
import Layout from "../components/layout"
import Seo from "../components/seo"


// <StaticImage
//     src="../images/example.png"
//     loading="eager"
//     width={64}
//     quality={95}
//     formats={["auto", "webp", "avif"]}
//     alt=""
//     style={{ marginBottom: `var(--space-3)` }}
// />

const IndexPage = () => (
    <main className="min-h-screen flex items-center">
        <div className="max-w-[1200px] w-full mx-auto p-8">
            <p className="text-3xl text-neutral-500">
                Hi,
            </p>
            <p className="text-xl">
                My name is Sam.
            </p>
        </div>
    </main>
)

/**
 * Head export to define metadata for the page
 *
 * See: https://www.gatsbyjs.com/docs/reference/built-in-components/gatsby-head/
 */
export const Head = () => {
    return (
        <>
            <link rel="preconnect" href="https://fonts.googleapis.com" />
            <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
            <link href="https://fonts.googleapis.com/css2?family=Space+Mono:ital,wght@0,400;0,700;1,400;1,700&display=swap" rel="stylesheet" />
            <Seo title="Home" />
            <html lang="en" className="text-black dark:text-white bg-uh-50 dark:bg-uh-950" />
            <body className="" />
        </>
    )
}

export default IndexPage
