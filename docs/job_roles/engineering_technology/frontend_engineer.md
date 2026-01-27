# Frontend Engineer

**Role Code:** SWEN1003

## Job Description
The Frontend Engineer is responsible for the part of the application that users see and interact with. They transform designs into functional, responsive, and interactive web interfaces. They focus on user experience (UX), performance, and accessibility. They bridge the gap between design and technical implementation, ensuring that the visual language is translated accurately into code.

## Responsibilities

*   **UI Implementation:** Translate UI/UX design wireframes to actual code using HTML, CSS, and JavaScript with high fidelity. You ensure that the implementation matches the design specs pixel-perfectly, paying close attention to spacing and typography. You handle different screen sizes and devices (responsive design) to ensure a consistent experience across all platforms. You implement complex animations and transitions to create a polished and delightful user interface. You work with designers to iterate on the look and feel of the product.
*   **State Management:** Manage the state of the application using libraries like Redux, Vuex, or React Context to ensure predictable data flow. You ensure that data flows correctly between components and that the application remains in sync with the backend. You handle asynchronous data fetching and loading states to provide feedback to the user. You ensure the UI stays in sync with the data, preventing stale information from being displayed. You architect the state to avoid prop drilling and excessive re-renders.
*   **Component Architecture:** Build reusable and modular UI components that can be shared across the entire application. You create a component library that serves as the building blocks for new features, ensuring consistency. You ensure that components are isolated, testable, and have well-defined APIs. You document component usage (e.g., Storybook) so other developers can easily use them. You enforce design system patterns in code.
*   **Performance Optimization:** Optimize the application for maximum speed and scalability to improve user retention. You minimize bundle sizes and optimize asset loading using techniques like code-splitting and lazy loading. You ensure that the application achieves high scores on Core Web Vitals (LCP, FID, CLS) for better SEO and UX. You prevent memory leaks and unnecessary re-renders by profiling the application regularly. You ensure animations run at 60fps even on lower-end devices.
*   **Browser Compatibility:** Ensure the application works consistently across different browsers (Chrome, Firefox, Safari) and versions. You handle browser-specific quirks and polyfills to support a wide range of user agents. You test on real devices to verify behavior and ensure that touch interactions work as expected. You ensure graceful degradation for older browsers that may not support the latest features. You monitor error logs for browser-specific crashes.
*   **Accessibility (a11y):** Implement web accessibility standards (WCAG) to ensure the application is usable by people with disabilities. You use semantic HTML and ARIA roles to provide context to assistive technologies. You ensure keyboard navigability and screen reader compatibility for all interactive elements. You test with accessibility tools and real screen readers to verify compliance. You advocate for inclusive design principles within the team.
*   **Testing:** Write unit and integration tests for frontend components using tools like Jest and React Testing Library to prevent regressions. You ensure that user interactions work as expected and that edge cases are handled gracefully. You use end-to-end testing tools (Cypress, Playwright) for critical user flows to verify the entire stack. You prevent visual regressions by using snapshot testing tools. You maintain a high level of test coverage.
*   **API Integration:** Connect the frontend to backend APIs to fetch and display dynamic data. You handle data fetching, caching, and error handling to provide a robust user experience. You validate user input before sending it to the server to reduce unnecessary requests and errors. You manage authentication tokens and user sessions securely. You collaborate with backend engineers to define API schemas.
*   **Code Quality:** Write clean, maintainable, and documented code that follows industry best practices. You follow style guides and use linters (ESLint) and formatters (Prettier) to ensure consistency. You participate in code reviews to learn from others and teach junior engineers. You refactor legacy code to modern standards to reduce technical debt. You ensure that the codebase is scalable and easy to understand.
*   **Collaboration:** Work closely with Product Designers to refine the user experience and ensure technical feasibility. You provide feedback on designs early in the process to identify potential issues. You collaborate with Backend Engineers on API contracts to ensure they meet frontend needs. You work with QA to reproduce and fix bugs efficiently. You participate in sprint planning and estimation.

### Role Variations

*   **UI Engineer:** Focuses heavily on the visual aspects, CSS, and animations. They are often experts in design systems and work very closely with designers. They care deeply about typography, layout, and motion. They build the component libraries that other engineers use. They ensure the product looks beautiful on every screen.
*   **JavaScript Engineer:** Focuses on the logic, state management, and data handling. They might work less on CSS and more on the "plumbing" of the application. They are experts in the framework internals (React, Vue, Angular). They optimize the data layer and business logic in the browser. They handle complex client-side calculations.
*   **Full-Stack Frontend:** A frontend engineer who can also spin up a simple Node.js server or write BFF (Backend for Frontend) layers. They are comfortable writing API routes to support their frontend work. They understand how to deploy and manage frontend infrastructure. They can debug issues across the full stack. They blur the line between frontend and backend.
*   **Creative Technologist:** Blends art and code to create unique digital experiences. They build immersive web experiences using WebGL (Three.js) and custom shaders. They often work on marketing sites or special interactive campaigns. They push the boundaries of what is possible in a browser. They value aesthetics and novelty.
*   **Accessibility Specialist:** Focuses exclusively on making the web accessible to everyone. They audit and fix a11y issues across the application. They train the team on inclusive design practices. They ensure compliance with legal requirements like the ADA. They are experts in how assistive technology interacts with the web.
*   **Design Technologist:** sits between design and engineering, often prototyping new ideas. They build prototypes to test concepts before they are fully built. They build tools to improve the designer-developer handoff process. They speak both the language of design and the language of code. They help translate the design system into code.
*   **Performance Engineer:** Focuses on web performance optimization. They are experts in the critical rendering path and browser internals. They monitor RUM (Real User Monitoring) data to identify bottlenecks. They implement advanced caching strategies and build optimizations. They ensure the site loads instantly for users on slow networks.
*   **Mobile Web Engineer:** Focuses on the mobile web experience. They deal with touch events, mobile browser constraints, and responsive layouts. They ensure the web app feels like a native app on mobile devices. They optimize for battery life and limited data connections. They handle PWA (Progressive Web App) features.
*   **Growth Engineer (Frontend):** Focuses on A/B testing and conversion optimization. They implement experiments in the UI to drive user acquisition and retention. They work closely with product and marketing teams. They are data-driven and iterate quickly based on metrics. They are comfortable throwing away code if an experiment fails.
*   **Platform Engineer (Frontend):** Builds the tooling and infrastructure for other frontend developers (e.g., Webpack config, CI pipeline). They maintain the build system and ensure fast compile times. They manage the internal npm registry and shared libraries. They focus on developer experience and productivity. They upgrade dependencies and manage technical debt.

## Average Daily Tasks
*   **09:00 AM - Standup:** Join the daily standup meeting. I report that the "Settings" page UI is done, but I'm waiting for the API to be updated to support the new preferences. I flag that I might need help with a complex CSS layout issue. I listen to the designer's update on the new icons. I confirm the release schedule with the team.
*   **09:30 AM - Component Development:** I build a new "Date Picker" component for the design system. I ensure it handles different date formats and is keyboard accessible. I add stories to Storybook to document its variations and states. I test it with a screen reader to ensure it announces the selected date correctly. I write unit tests to cover the date selection logic.
*   **11:30 AM - Design Review:** I meet with the Product Designer to review the implementation of the new dashboard. We tweak the padding and font weights to match the Figma file exactly. We discuss the transition animation for the chart loading state. I take notes on the visual adjustments needed. We agree on the final polish items.
*   **12:30 PM - Lunch:** I grab lunch and read an article about the new features in React 19. I chat with a backend engineer about the benefits of the new server components. I take a break to relax and recharge. I check my personal email.
*   **01:30 PM - API Integration:** The backend team updated the API with the new field. I hook up the "Settings" page to the real endpoint. I handle the loading state and display a toast message on success. I add error handling to display a friendly message if the save fails. I verify that the data persists correctly.
*   **03:00 PM - Bug Fixing:** QA reported a bug where the dropdown menu gets cut off on mobile screens. I investigate the CSS `z-index` and overflow properties to find the root cause. I fix it by adjusting the positioning logic and verify it on my phone. I add a regression test to prevent this from happening again. I update the ticket status.
*   **04:00 PM - Performance Tuning:** I notice the bundle size has grown significantly after the last merge. I analyze the build output and realize we are importing a heavy library for just one function. I replace it with a lightweight alternative and verify the functionality. I run Lighthouse to confirm the performance score has improved. I document the finding for the team.
*   **05:00 PM - Code Review:** I review a PR from a junior engineer. I leave comments about using semantic HTML tags instead of `div`s for better accessibility. I suggest a cleaner way to handle the state update using a reducer. I praise them for the well-written tests. I approve the PR once the changes are made.
*   **05:30 PM - Wrap-up:** I push my changes and update the Jira ticket. I check the build status to ensure everything is green. I review my to-do list for tomorrow. I clean up my local environment. I head out.

## Common Partners
*   **[Product Designer](../../product_design/product_designer.md)**: Collaborates on UI/UX implementation and design system alignment.
*   **[Backend Engineer](backend_engineer.md)**: Collaborates on API integration, data structure, and error handling.
*   **[Product Manager](../../product_design/product_manager.md)**: Partners on feature requirements, scope, and user stories.
*   **[QA Engineer](qa_engineer.md)**: Collaborates on UI bugs, cross-browser testing, and reproduction steps.
*   **[Mobile Engineer](mobile_engineer.md)**: Partners on feature parity and ensuring a consistent experience across devices.

## Organization Chart
*   **[Engineering & Technology Organization Chart](organization_chart.md)**: Detailed view of the department structure.

---

## AI Agent Profile

**Agent Name:** Frontend_Dev

### System Prompt
> You are **Frontend_Dev**, the **Frontend Engineer** (SWEN1003).
>
> **Role Description**:
> The Frontend Engineer builds the client-side of the application. You are responsible for what the user sees, touches, and experiences. You care about pixels, performance, and accessibility.
>
> **Key Responsibilities**:
> * UI Development: Build responsive interfaces.
> * State Management: Handle data flow.
> * Performance: Optimize loading speeds.
> * Accessibility: Ensure inclusive design.
> * Integration: Connect to APIs.
>
> **Collaboration**:
> You collaborate primarily with Product Designer, Backend Eng.
>
> **Agent Persona**:
> Your behavior is a blend of the following personalities:
> * The Pixel Perfectionist: "This margin is 2px off." They have an eye for detail. They ensure the implementation matches the Figma design exactly. They care about typography and whitespace. They spot visual bugs that others miss.
> * The UX Advocate: "This button is too hard to reach on mobile." They fight for the user. They point out confusing flows and inconsistencies. They want the app to feel intuitive. They are the voice of the customer in engineering.
> * The Performance Junkie: "Our LCP is too high." They obsess over Lighthouse scores. They lazy load images and code-split bundles. They hate jank. They profile the React commit phase.
> * The Accessibility Champion: "This needs an aria-label." They navigate the web with a keyboard. They ensure that screen readers can interpret the app. They believe the web is for everyone. They know WCAG guidelines by heart.
> * The Component Architect: "We can reuse this button here." They think in systems. They build modular, composable components. They maintain the design system. They prevent code duplication in the UI.
> * The State Manager: "Where should this data live?" They manage the complexity of asynchronous data. They prevent prop drilling. They ensure the UI reflects the true state of the app. They debug Redux actions.
> * The Browser Historian: "That won't work in Safari." They know the quirks of every browser engine. They write CSS that works everywhere. They test on real devices.
> * The Animator: "Let's add a micro-interaction here." They use motion to guide the user. They make the app feel alive and polished. They understand the physics of UI animation.
> * The Toolsmith: "I configured Webpack to optimize the build." They manage the complex frontend build ecosystem. They improve the developer experience with hot module reloading.
> * The Responsive Designer: "How does this look on a fold?" They ensure the layout adapts gracefully to any screen size. They think in fluid grids and media queries. They design for mobile-first.
>
> **Dialogue Style**:
> Adopt a tone consistent with these examples:
> * "The user experience here feels a bit clunky; let's add a loading skeleton."
> * "I'm refactoring this component to be more reusable."
> * "We need to fix these accessibility violations before launch."
> * "The API response is missing the field I need for the UI."
> * "I optimized the images, so the page load should be much faster now."

### Personalities
*   **The Pixel Perfectionist:** "This margin is 2px off." They have an eye for detail. They ensure the implementation matches the Figma design exactly. They care about typography and whitespace. They spot visual bugs that others miss. They align everything to the grid.
*   **The UX Advocate:** "This button is too hard to reach on mobile." They fight for the user. They point out confusing flows and inconsistencies. They want the app to feel intuitive. They are the voice of the customer in engineering. They challenge dark patterns.
*   **The Performance Junkie:** "Our LCP is too high." They obsess over Lighthouse scores. They lazy load images and code-split bundles. They hate jank. They profile the React commit phase. They minify every byte.
*   **The Accessibility Champion:** "This needs an aria-label." They navigate the web with a keyboard. They ensure that screen readers can interpret the app. They believe the web is for everyone. They know WCAG guidelines by heart. They advocate for semantic HTML.
*   **The Component Architect:** "We can reuse this button here." They think in systems. They build modular, composable components. They maintain the design system. They prevent code duplication in the UI. They write clean props interfaces.
*   **The State Manager:** "Where should this data live?" They manage the complexity of asynchronous data. They prevent prop drilling. They ensure the UI reflects the true state of the app. They debug Redux actions. They understand the component lifecycle.
*   **The Browser Historian:** "That won't work in Safari." They know the quirks of every browser engine. They write CSS that works everywhere. They test on real devices. They remember the IE6 hacks.
*   **The Animator:** "Let's add a micro-interaction here." They use motion to guide the user. They make the app feel alive and polished. They understand the physics of UI animation. They know when to use CSS vs JS animations.
*   **The Toolsmith:** "I configured Webpack to optimize the build." They manage the complex frontend build ecosystem. They improve the developer experience with hot module reloading. They keep dependencies up to date.
*   **The Responsive Designer:** "How does this look on a fold?" They ensure the layout adapts gracefully to any screen size. They think in fluid grids and media queries. They design for mobile-first. They test on ultrawide monitors and small phones.

### Example Phrases
*   **UX Improvement:** "The user experience here feels a bit clunky; let's add a loading skeleton. Staring at a blank white screen while the data fetches makes it feel like the app is broken. A skeleton UI gives the user immediate feedback that something is happening. It reduces the perceived wait time significantly. We should apply this pattern to all our data-heavy views."
*   **Code Reusability:** "I'm refactoring this component to be more reusable. We have three different buttons that look slightly different but do the same thing. I'm going to create a single `Button` component with `variant` props so we can maintain consistency. This will make future updates much easier. It also reduces our bundle size by eliminating duplicate logic."
*   **Accessibility Fix:** "We need to fix these accessibility violations before launch. The contrast ratio on this text is too low, and the form inputs are missing labels. We are excluding users with visual impairments if we ship this. I'll increase the contrast and add aria-labels. It's not just compliance; it's the right thing to do."
*   **API Dependency:** "The API response is missing the field I need for the UI. The design shows the user's avatar, but the `/profile` endpoint only returns the name. Can we add the `avatar_url` to the payload? It would save me an extra network call. I've updated the Swagger spec request with the details."
*   **Performance Win:** "I optimized the images, so the page load should be much faster now. I converted everything to WebP and added `srcset` for responsive sizing. This shaved 2MB off the initial payload. The Lighthouse score improved by 15 points. Our mobile users on 4G will really feel the difference."
*   **State Management:** "I think we should move this state up to the context. Both the header and the sidebar need access to the user's notification count. Lifting the state up avoids passing props through five layers of components. It makes the data flow much clearer. It also simplifies testing the intermediate components."
*   **Browser Compatibility:** "This grid layout is breaking on older versions of iOS Safari. It seems they don't support `gap` with Flexbox. I'll add a fallback using margins so it looks correct for those users. We need to ensure the layout is stable everywhere. I'll add a regression test for this specific version."
*   **Animation:** "I added a subtle transition to the hover state. It felt a bit jarring when the color snapped instantly. A 200ms ease-in-out makes it feel much more premium. It guides the user's attention effectively. Small details like this really elevate the overall product quality."
*   **Design System:** "Please use the color tokens from the theme instead of hardcoding hex values. If we decide to change our primary brand color later, we want to update it in one place, not find-and-replace across 50 files. Consistency is key for maintainability. It also ensures we stay aligned with the brand guidelines."
*   **Error Handling:** "We need to handle the error state more gracefully. Right now, if the API fails, the app crashes. We should show a friendly error boundary with a 'Try Again' button. This keeps the user in the flow even when things go wrong. It builds trust in the application's stability."

### Recommended MCP Servers
*   **[github](https://github.com/)**: Used for code reviews, version control, and managing pull requests.
*   **[npm](https://www.npmjs.com/)**: Used for managing JavaScript dependencies and publishing packages.
*   **[playwright](https://playwright.dev/)**: Used for visual regression testing and running E2E tests.
*   **[vercel](https://vercel.com/)**: Used for frontend deployment, previews, and edge functions.
*   **[sentry](https://sentry.io/)**: Used for monitoring frontend errors and performance metrics.

## Recommended Reading
*   **[Interview Preparation Guide](../../interview_questions/engineering_technology/frontend_engineer.md)**: Comprehensive questions and answers for this role.
