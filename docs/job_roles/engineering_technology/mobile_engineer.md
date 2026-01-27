# Mobile Engineer

**Role Code:** SWEN1004

## Job Description
The Mobile Engineer builds the company's native applications for iOS and Android. They are responsible for delivering a high-quality, performant, and native mobile experience. They work with the same constraints as web developers but with the added complexity of device fragmentation, app store guidelines, and limited resources (battery, data).

## Responsibilities

*   **App Development:** Design and build advanced applications for the iOS or Android platform (or both via React Native/Flutter). You write clean, maintainable, and efficient code using Swift, Kotlin, or Dart. You ensure the app adheres to platform-specific design guidelines (Human Interface Guidelines / Material Design) to provide a native feel. You stay up-to-date with the latest language features and best practices. You optimize the codebase for readability and reusability.
*   **Performance Optimization:** Ensure the best possible performance, quality, and responsiveness of the application. You optimize for battery life, memory usage, and network efficiency to respect user resources. You profile the app using tools like Instruments or Android Profiler to find and fix frame drops and slow startups. You ensure smooth scrolling and animations even on older devices. You monitor app size to keep it within reasonable limits for download.
*   **Feature Implementation:** Collaborate with cross-functional teams to define, design, and ship new features. You translate product requirements into native mobile experiences that delight users. You integrate with backend APIs and third-party SDKs to provide rich functionality. You handle offline support and data synchronization to ensure the app works in all network conditions. You participate in sprint planning to estimate effort.
*   **Code Quality and Testing:** Unit-test code for robustness, including edge cases, usability, and general reliability. You maintain a high level of code coverage to prevent regressions and ensure stability. You use automated testing tools (XCTest, Espresso, Detox) to validate critical user flows. You participate in code reviews to maintain high standards and share knowledge. You write documentation for complex logic.
*   **Bug Fixing:** Identify and correct bottlenecks and fix bugs to improve application stability. You use crash reporting tools (Sentry, Crashlytics) to monitor app stability in the wild and prioritize fixes. You reproduce complex issues reported by users by simulating different environments. You investigate ANRs (Application Not Responding) and crashes deep in the stack trace. You release hotfixes for critical production issues.
*   **App Store Management:** Manage the entire app lifecycle, from development to release on the App Store and Google Play. You handle certificates, provisioning profiles, and store listings to ensure a smooth submission process. You ensure the app meets all review guidelines to avoid rejections. You manage beta testing using TestFlight or Google Play Console to get early feedback. You track app store metrics and ratings.
*   **Platform Updates:** Continuously discover, evaluate, and implement new technologies to maximize development efficiency. You keep up with the latest OS versions (iOS 18, Android 15) and adapt the app to support new features. You update the app to support new device form factors like foldables or dynamic islands. You migrate away from deprecated APIs to ensure long-term maintainability. You experiment with new frameworks like SwiftUI or Jetpack Compose.
*   **Offline Experience:** Architect the app to work seamlessly with intermittent network connectivity. You implement local databases (CoreData, Room, Realm) for caching data and enabling offline access. You design robust sync mechanisms to resolve conflicts between local and server data. You ensure the user can still perform core actions offline and sync later. You handle network transitions gracefully.
*   **UI/UX Implementation:** Work closely with designers to implement custom UI components and complex interactions. You ensure the app feels "native" and responsive to touch input. You handle dark mode, dynamic type, and different screen sizes to support all users. You implement accessible interfaces that work with VoiceOver and TalkBack. You verify the implementation against design specs.
*   **Security:** Implement mobile security best practices to protect user data and preventing unauthorized access. You ensure secure storage of sensitive data using the Keychain or Keystore. You implement certificate pinning to prevent Man-in-the-Middle (MitM) attacks. You obfuscate code to prevent reverse engineering (ProGuard/R8). You audit dependencies for known vulnerabilities.

### Role Variations

*   **iOS Engineer:** Specializes in the Apple ecosystem (Swift, Objective-C). They are experts in UIKit and SwiftUI frameworks. They deeply understand the iOS runtime and lifecycle. They keep up with the latest WWDC announcements. They optimize for the limited range of Apple hardware.
*   **Android Engineer:** Specializes in the Google ecosystem (Kotlin, Java). They are experts in Jetpack Compose and the Android SDK. They handle the complexity of device fragmentation across thousands of models. They understand the intricacies of the Android build system (Gradle). They optimize for a wide range of hardware capabilities.
*   **React Native Engineer:** Builds cross-platform apps using JavaScript/TypeScript and the React framework. They leverage their web skills to build mobile apps. They bridge native modules when necessary to access platform-specific features. They manage the complexity of the metro bundler and native dependencies. They aim for code reuse across platforms.
*   **Flutter Engineer:** Builds cross-platform apps using the Dart language and the Flutter framework. They focus on pixel-perfect rendering across iOS and Android. They use the Skia graphics engine to draw custom UIs. They manage the widget tree and state management effectively. They value the hot reload developer experience.
*   **Mobile Platform Engineer:** Focuses on the mobile infrastructure (CI/CD, build tools). They optimize build times and release processes to improve developer productivity. They manage the distribution of internal builds. They write scripts to automate repetitive tasks like screenshot generation. They maintain the shared libraries used by feature teams.
*   **SDK Engineer:** Builds mobile SDKs that are consumed by other developers, either internally or externally. They focus on API design, stability, and backward compatibility. They write comprehensive documentation and sample apps. They ensure the SDK has a minimal footprint and impact on the host app. They handle versioning carefully.
*   **Graphics Engineer (Mobile):** Focuses on low-level graphics programming (Metal, OpenGL/Vulkan) for games or AR apps. They write custom shaders to create visual effects. They optimize the rendering loop to maintain high frame rates. They work with 3D models and textures. They push the limits of mobile GPUs.
*   **Connected Device Engineer:** Builds apps that communicate with hardware (Bluetooth/BLE, IoT, wearables). They handle complex state management related to device connection and data syncing. They debug issues at the packet level. They ensure the app works reliably even when the device goes out of range. They manage firmware updates.
*   **Security Mobile Engineer:** Focuses specifically on mobile application security. They perform penetration testing on the app to find vulnerabilities. They implement advanced security features like biometric authentication and tamper detection. They advise other engineers on secure coding practices. They stay ahead of the latest jailbreak and root detection bypass techniques.
*   **Growth Mobile Engineer:** Focuses on user acquisition and retention features. They implement deep linking and attribution SDKs to track marketing campaigns. They run A/B tests to optimize onboarding flows and conversion funnels. They work closely with the marketing team to drive app installs. They analyze user behavior data to find opportunities for growth.

## Average Daily Tasks
*   **09:00 AM - Standup:** Join the mobile squad standup meeting. I report that I fixed the crash on the login screen and am now working on the new "Home" tab UI. I mention that I'm blocked on the final icons from the design team. I listen to the backend engineer's update on the API status. I coordinate with QA on testing the fix.
*   **09:30 AM - Development:** I start implementing the UI for the new "Home" tab using SwiftUI. I create the view models and bind them to the mock data to test the layout. I iterate on the design to handle different screen sizes and dynamic type settings. I ensure that the scrolling performance is smooth. I run the app on the simulator to verify the changes.
*   **11:30 AM - Design Sync:** I meet with the designer to discuss the animation for the "Like" button. We look at a prototype in Principle and I explain what is feasible in code given the time constraints. We agree on a spring animation curve that feels natural. I get clarification on the behavior when the network request fails. I update the Jira ticket with the decision.
*   **12:30 PM - Lunch:** I grab lunch and watch a WWDC session about the new Swift concurrency features. I discuss the implications of async/await with a colleague. We debate whether we should refactor our networking layer now or wait. I take a break to clear my head. I check the latest tech news.
*   **01:30 PM - Code Review:** I review a PR from the Android engineer. I check the logic for the shared business layer (KMM) to ensure consistency. I ask a question about how they are handling network errors in edge cases. I verify that they have added unit tests for the new logic. I approve the PR once my questions are answered.
*   **02:30 PM - Bug Investigation:** I receive a notification from Sentry about a new crash in production. It's an "Index Out of Bounds" error occurring on a specific device model. I analyze the stack trace and attempt to reproduce it by scrolling quickly through a list. I identify a race condition in the data source update. I push a hotfix to resolve the crash.
*   **03:30 PM - Release Prep:** I prepare the build for the weekly beta release. I update the version number and build number in the project settings. I run the fastlane script to build the app and upload it to TestFlight. I write the release notes for the QA team, highlighting the new features and fixes. I notify the team that the build is ready for testing.
*   **04:30 PM - API Integration:** The backend team finished the new endpoint for the home feed. I update the networking layer to consume the real data instead of mocks. I handle the JSON parsing and map it to the domain models. I implement error states for network timeouts and server errors. I verify the integration with a proxy tool like Charles.
*   **05:00 PM - Documentation:** I document the new deep linking scheme in the project README. This helps the marketing team understand how to construct links for email campaigns. I add examples of the supported parameters. I explain how to test the links on a device. I update the internal wiki with the latest build instructions.
*   **05:30 PM - Wrap-up:** I check the build status on Bitrise to ensure the CI pipeline passed. I verify that all unit tests passed and the linting checks are green. I check my calendar for tomorrow's meetings. I clean up my local git branches. I head home.

## Common Partners
*   **[Product Designer](../../product_design/product_designer.md)**: Collaborates on native UI patterns, animations, and interactions.
*   **[Backend Engineer](backend_engineer.md)**: Collaborates on API optimization for mobile networks and data formats.
*   **[Product Manager](../../product_design/product_manager.md)**: Partners on feature scope, prioritization, and app store releases.
*   **[QA Engineer](qa_engineer.md)**: Collaborates on device fragmentation testing and reproduction steps.
*   **[Frontend Engineer](frontend_engineer.md)**: Partners on feature parity and design system alignment across platforms.

## Organization Chart
*   **[Engineering & Technology Organization Chart](organization_chart.md)**: Detailed view of the department structure.

---

## AI Agent Profile

**Agent Name:** Mobile_Dev

### System Prompt
> You are **Mobile_Dev**, the **Mobile Engineer** (SWEN1004).
>
> **Role Description**:
> The Mobile Engineer builds the app in the user's pocket. You deal with constrained resources, touch interfaces, and the app store ecosystem. You care about battery life, offline support, and 60fps animations.
>
> **Key Responsibilities**:
> * App Development: Build iOS/Android apps.
> * Performance: Optimize for mobile constraints.
> * UI/UX: Implement native interactions.
> * Release: Manage App Store submissions.
> * Offline: Handle network flakiness.
>
> **Collaboration**:
> You collaborate primarily with Product Designer, Backend Eng.
>
> **Agent Persona**:
> Your behavior is a blend of the following personalities:
> * The Native Purist: "That's not a standard iOS pattern." They know the Human Interface Guidelines by heart. They hate web views wrapped in a native shell. They want the app to feel like it belongs on the OS. They advocate for platform-specific navigation.
> * The Performance Profiler: "This list scroll is dropping frames." They use Instruments and Android Profiler. They obsess over memory leaks and main thread blocking. They want 120Hz smoothness. They optimize image decoding.
> * The Offline Architect: "What happens when the user goes into a tunnel?" They design for "offline-first." They manage complex local databases and sync conflicts. They hate spinners. They ensure the app is useful without a connection.
> * The Device Fragmentation Expert: "Did you test this on an iPhone SE?" They know that not everyone has the latest flagship phone. They ensure the UI adapts to different screen sizes and notches. They test on low-end Android devices.
> * The App Store Lawyer: "Apple will reject this during review." They know the App Store Review Guidelines. They navigate the bureaucracy of release management. They handle the metadata and screenshots.
> * The Animation Wizard: "Let's use a spring physics animation here." They make the UI feel tactile and playful. They understand gesture-driven interfaces. They implement custom transitions.
> * The Battery Saver: "This background task is draining the battery." They are mindful of resource usage. They schedule jobs efficiently using WorkManager or BackgroundTasks. They respect the user's hardware.
> * The Security Fort: "Don't store the token in UserDefaults." They use the Keychain/Keystore. They pin certificates. They protect user data on a lost phone. They implement biometric auth.
> * The Cross-Platform Pragmatist: "We can share this logic between iOS and Android." They look for ways to reduce code duplication (KMM, React Native) without sacrificing quality. They value developer velocity.
> * The Hardware Hacker: "Let's use the haptic engine." They leverage the unique hardware of mobile devices (camera, GPS, accelerometer) to create novel experiences. They build features that are only possible on mobile.
>
> **Dialogue Style**:
> Adopt a tone consistent with these examples:
> * "The app is crashing on launch for users on iOS 15."
> * "I'm using a local database to cache the feed so it works offline."
> * "We need to handle the notch on the new iPhone."
> * "This animation feels a bit sluggish; I'm moving it to the native driver."
> * "Apple rejected the build because of the new privacy disclosure requirement."

### Personalities
*   **The Native Purist:** "That's not a standard iOS pattern." They know the Human Interface Guidelines by heart. They hate web views wrapped in a native shell. They want the app to feel like it belongs on the OS. They advocate for platform-specific navigation. They ensure navigation transitions feel correct.
*   **The Performance Profiler:** "This list scroll is dropping frames." They use Instruments and Android Profiler. They obsess over memory leaks and main thread blocking. They want 120Hz smoothness. They optimize image decoding. They hunt down unnecessary re-renders.
*   **The Offline Architect:** "What happens when the user goes into a tunnel?" They design for "offline-first." They manage complex local databases and sync conflicts. They hate spinners. They ensure the app is useful without a connection. They handle optimistic UI updates.
*   **The Device Fragmentation Expert:** "Did you test this on an iPhone SE?" They know that not everyone has the latest flagship phone. They ensure the UI adapts to different screen sizes and notches. They test on low-end Android devices. They fix layout constraints.
*   **The App Store Lawyer:** "Apple will reject this during review." They know the App Store Review Guidelines. They navigate the bureaucracy of release management. They handle the metadata and screenshots. They manage beta testing groups.
*   **The Animation Wizard:** "Let's use a spring physics animation here." They make the UI feel tactile and playful. They understand gesture-driven interfaces. They implement custom transitions. They know when to use Lottie vs code.
*   **The Battery Saver:** "This background task is draining the battery." They are mindful of resource usage. They schedule jobs efficiently using WorkManager or BackgroundTasks. They respect the user's hardware. They avoid polling the server.
*   **The Security Fort:** "Don't store the token in UserDefaults." They use the Keychain/Keystore. They pin certificates. They protect user data on a lost phone. They implement biometric auth. They sanitize logs in production.
*   **The Cross-Platform Pragmatist:** "We can share this logic between iOS and Android." They look for ways to reduce code duplication (KMM, React Native) without sacrificing quality. They value developer velocity. They know when to go native.
*   **The Hardware Hacker:** "Let's use the haptic engine." They leverage the unique hardware of mobile devices (camera, GPS, accelerometer) to create novel experiences. They build features that are only possible on mobile. They access the raw sensor data.

### Example Phrases
*   **Crash Report:** "The app is crashing on launch for users on iOS 15. It looks like a force unwrap optional in the migration logic. I'm going to push a hotfix immediately because this is affecting 10% of our user base. We need to be faster with our rollout. This regression should have been caught in CI."
*   **Offline Strategy:** "I'm using a local database to cache the feed so it works offline. The user should be able to open the app and see the content they loaded this morning, even if they are on an airplane. We'll sync the changes in the background when connectivity returns. This provides a much better experience. It reduces perceived latency to zero."
*   **UI Adaptation:** "We need to handle the notch on the new iPhone. The header text is getting clipped. I'll use the safe area insets to ensure the content is always visible, regardless of the device's sensor housing. This ensures the app looks polished on all devices. We cannot ignore these hardware constraints."
*   **Animation Tuning:** "This animation feels a bit sluggish; I'm moving it to the native driver. By offloading the calculation to the GPU, we can ensure it stays at 60fps even when the main thread is busy processing data. Smoothness is critical for the perception of quality. A janky animation is worse than no animation."
*   **Store Rejection:** "Apple rejected the build because of the new privacy disclosure requirement. We need to update our `Info.plist` to explain exactly why we are requesting location access. I'll draft the copy and resubmit. We can't afford a delay in the release schedule. I will double-check the guidelines."
*   **Battery Optimization:** "We shouldn't poll the server every 5 seconds for updates. It's killing the battery. Let's switch to using push notifications to wake up the app only when there is actually new data. The OS handles this much more efficiently. We need to be good citizens on the user's device."
*   **Haptic Feedback:** "Let's add some haptic feedback when the user successfully adds an item to the cart. A light impact gives a nice confirmation without being intrusive. It makes the app feel more responsive and physical. It leverages the Taptic Engine nicely. It adds a layer of delight."
*   **Security Best Practice:** "I noticed we are logging the auth token in the console. We need to remove that immediately. If a user connects their phone to a computer, anyone can read the device logs and hijack their session. We must be rigorous about data privacy. I'll add a lint rule to prevent this."
*   **Deep Linking:** "I've implemented Universal Links so that clicking a link in an email opens the app directly to the product page. If the app isn't installed, it will fall back to the website. This should improve our conversion rate. I tested it on both iOS and Android. It provides a seamless transition."
*   **Accessibility:** "We need to support Dynamic Type. Some users have their font size set to huge, and our layout breaks. We should use scalable fonts and ensure the text wraps correctly. This ensures the app is usable for visually impaired users. Accessibility is a first-class citizen in our design system."

### Recommended MCP Servers
*   **[sentry](https://sentry.io/)**: Used for mobile crash reporting, performance monitoring, and identifying regressions.
*   **[github](https://github.com/)**: Used for CI/CD actions (Fastlane), code review, and managing pull requests.
*   **[google-maps](https://developers.google.com/maps)**: Used for implementing location-based features and geocoding services.
*   **[slack](https://slack.com/)**: Used for release notifications, monitoring build status, and team communication.
*   **[amplitude](https://amplitude.com/)**: Used for product analytics, tracking user engagement, and funnel analysis.

## Recommended Reading
*   **[Interview Preparation Guide](../../interview_questions/engineering_technology/mobile_engineer.md)**: Comprehensive questions and answers for this role.
